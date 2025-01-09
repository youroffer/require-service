package parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/himmel520/uoffer/require/config"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/rs/zerolog/log"
)

type Parser struct {
	analyticRepo AnalyticRepo
	filterRepo   FilterRepo
	cache        Cache
	qe           DBTX
	cfg          config.API_HH

	apiBaseURL string
	filter     map[string]struct{}

	limiter *time.Ticker
	mu      sync.RWMutex
	wg      sync.WaitGroup
}

type (
	AnalyticRepo interface {
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.AnalyticResp, error)
		Update(ctx context.Context, qe repository.Querier, id int, analytic *entity.AnalyticUpdate) (*entity.AnalyticResp, error)
	}

	FilterRepo interface {
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Filter, error)
	}

	Cache interface {
		Set(ctx context.Context, key string, value any) error
	}

	DBTX interface {
		DB() repository.Querier
	}
)

type ParserParams struct {
	Cfg          config.API_HH
	AnalyticRepo AnalyticRepo
	FilterRepo   FilterRepo
	Cache        Cache
	DBTX         DBTX
}

func NewParser(params ParserParams) *Parser {
	return &Parser{
		cfg:          params.Cfg,
		analyticRepo: params.AnalyticRepo,
		filterRepo:   params.FilterRepo,
		cache:        params.Cache,
		qe:           params.DBTX,

		apiBaseURL: "https://api.hh.ru/vacancies",
		// ограничение запросов к апи hh.ru - 30 в сек
		limiter: time.NewTicker(time.Second / 30),
		filter:  map[string]struct{}{},
	}

}

func (p *Parser) SetHeaders(req *http.Request) {
	req.Header.Set("User-Agent", fmt.Sprintf("%s (%s)", p.cfg.AppName, p.cfg.Email))
	req.Header.Set("Authorization", p.cfg.AccessToken)
}

func (p *Parser) Parse(ctx context.Context) {
	// обновляем фильтр
	p.mustFilter()

	// получаем аналитику
	analytics, err := p.analyticRepo.Get(ctx, p.qe.DB(), repository.PaginationParams{})
	if err != nil {
		log.Err(err).Msg("get analytics")
		return
	}

	log.Info().Int("queries", len(analytics)).Msg("Start parsing")
	for _, analytic := range analytics {
		select {
		case <-ctx.Done():
			p.wg.Wait()
			p.limiter.Stop()
			return
		default:
			p.wg.Add(1)
			go func() {
				defer p.wg.Done()

				params := fmt.Sprintf("?area=113&text=%s&per_page=100", analytic.SearchQuery)
				positionUrl := p.apiBaseURL + params

				log.Debug().
					Str("position", analytic.PostTitle).
					Str("positionUrl", positionUrl).
					Str("searchQuery", analytic.SearchQuery).Msg("Parsing analytic")

				// получаем все страницы
				allPages, err := p.fetchPagesWithVacancies(ctx, positionUrl)
				if err != nil {
					var pageErr *pageFetchError
					if errors.As(err, &pageErr) {
						log.Err(pageErr.Err).Int("page", pageErr.Page).Msg("get pages")
						return
					}

					log.Err(err).Str("position", analytic.PostTitle).Msg("get pages")
					return
				}
				analytic.VacanciesNum = entity.NewOptional(allPages.Found)

				log.Debug().
					Str("position", analytic.PostTitle).
					Int("pages", allPages.Pages).
					Int("Found", allPages.Found).
					Msg("Getting vacancies data")

				vacancies, err := p.fetchVacancies(ctx, allPages.Items)
				if err != nil {
					var pageErr *vacancyFetchError
					if errors.As(err, &pageErr) {
						log.Err(pageErr.Err).
							Str("position", analytic.PostTitle).
							Str("vacancyID", pageErr.VacancyID)
						return
					}

					log.Err(err).Str("position", analytic.PostTitle).Msg("get vacancies")
					return
				}

				// получаем 100 самых встречаемых ключевых слов
				keywords := p.mustTopKeywords(vacancies, 100)

				// получаем 100 самых встречаемых навыков
				skills := p.mustTopSkills(vacancies, 100)

				// фиксируем время обновления
				analytic.ParseAt = entity.NewOptional(time.Now())
				// подготовка структуры для кэша
				analyticWords := &entity.AnalyticWithWords{
					Analytic: &entity.AnalyticResp{
						ID:           analytic.ID,
						SearchQuery:  analytic.SearchQuery,
						ParseAt:      analytic.ParseAt,
						VacanciesNum: analytic.VacanciesNum,
					},
					Skills:   skills,
					Keywords: keywords,
				}

				// отправляем в кэш
				if err := p.cache.Set(context.Background(), fmt.Sprintf(cache.AnalyticKeyFmt, analytic.ID), analyticWords); err != nil {
					log.Err(err).Str("position", analytic.PostTitle).Msg("set cache")
					return
				}

				if _, err := p.analyticRepo.Update(ctx, p.qe.DB(), analytic.ID, &entity.AnalyticUpdate{
					VacanciesNum: analytic.VacanciesNum,
					ParseAt:      analytic.ParseAt}); err != nil {
					log.Err(err).Str("position", analytic.PostTitle).Msg("update analytic")
				}
			}()
		}
	}
}

func (p *Parser) mustFilter() {
	filterMap := map[string]struct{}{}

	filters, err := p.filterRepo.Get(context.Background(), p.qe.DB(), repository.PaginationParams{})
	if err != nil {
		log.Err(err).Msg("get filters")
	}

	// преобразуем список в мапу для быстрого поиска
	for _, filter := range filters {
		filterMap[filter.Word] = struct{}{}
	}

	p.filter = filterMap
}

// Получаем количество страниц и вакансий, а также все id вакансий
func (p *Parser) fetchPagesWithVacancies(ctx context.Context, positionUrl string) (*PageResp, error) {
	// ограничение - один запрос в 33ms
	<-p.limiter.C

	pagesResp, err := p.fetchPage(positionUrl)
	if err != nil {
		return nil, &pageFetchError{Page: 0, Err: err}
	}

	for i := 1; i < pagesResp.Pages; i++ {
		if err := ctx.Err(); err != nil {
			return pagesResp, err
		}

		// ограничение - один запрос в 33ms
		<-p.limiter.C

		page, err := p.fetchPage(fmt.Sprintf("%s&page=%d", positionUrl, i))
		if err != nil {
			return nil, &pageFetchError{Page: i, Err: err}
		}

		pagesResp.Items = append(pagesResp.Items, page.Items...)

	}

	return pagesResp, nil
}

func (p *Parser) fetchPage(pageUrl string) (*PageResp, error) {
	req, err := http.NewRequest(http.MethodGet, pageUrl, nil)
	if err != nil {
		return nil, err
	}
	p.SetHeaders(req)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch page: status %v, response: %s", resp.StatusCode, string(body))
	}

	var page *PageResp
	if err := json.Unmarshal(body, &page); err != nil {
		return nil, err
	}

	return page, nil
}

func (p *Parser) fetchVacancies(ctx context.Context, vacancyIDs []VacancyID) ([]*Vacancy, error) {
	vacancies := []*Vacancy{}
	// идем по каждой вакансии
	for _, vacancyID := range vacancyIDs {
		if err := ctx.Err(); err != nil {
			return vacancies, err
		}

		// ограничение - один запрос в 33ms
		<-p.limiter.C

		// формируем ссылку на вакансию
		vacancyUrl := fmt.Sprintf("%v/%v", p.apiBaseURL, vacancyID.ID)

		req, err := http.NewRequest(http.MethodGet, vacancyUrl, nil)
		if err != nil {
			return nil, err
		}
		p.SetHeaders(req)

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode != http.StatusOK {
			return nil, &vacancyFetchError{
				VacancyID: vacancyID.ID,
				Err:       fmt.Errorf("failed to fetch vacancy: status %v, response: %s", resp.StatusCode, string(body))}
		}

		var vacancy *Vacancy
		if err := json.Unmarshal(body, &vacancy); err != nil {
			return nil, err
		}

		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (p *Parser) mustTopSkills(vacancies []*Vacancy, limit int) []*entity.TopWords {
	skills := map[string]int{}

	// считаем количество ключевых слов в каждой вакансии
	for _, vacancy := range vacancies {
		for _, skill := range vacancy.Skills {
			if _, ok := skills[skill.Name]; !ok {
				// добавляем скил
				skills[skill.Name] = 1
				continue
			}

			// обновляем счетчик
			skills[skill.Name] += 1
		}
	}

	return getTopWords(skills, limit)
}

func (p *Parser) mustTopKeywords(vacancies []*Vacancy, limit int) []*entity.TopWords {
	keywords := map[string]int{}

	// считаем количество ключевых слов в каждой вакансии
	for _, vacancy := range vacancies {
		// отчистить строку от html тэгов и привести к нижнему регистру
		text := strings.ToLower(stripHTMLTags(vacancy.Description))

		// получаем список всех слов
		words := extractEnglishWords(text)

		// считаем слова
		for _, word := range words {
			if _, ok := keywords[word]; !ok {
				p.mu.RLock()
				if _, ok := p.filter[word]; ok {
					p.mu.RUnlock()
					// слово находится в фильтре
					continue
				}
				p.mu.RUnlock()

				// добавляем слово
				if !isNumber(word) {
					keywords[word] = 1
				}

				continue
			}

			// обновляем счетчик
			keywords[word] += 1
		}
	}

	return getTopWords(keywords, limit)
}

// stripHTMLTags отчищает строку от html тэгов
func stripHTMLTags(s string) string {
	re := regexp.MustCompile(`<[^>]*>`)
	return re.ReplaceAllString(s, "")
}

// extractEnglishWords извлекает все английские слова и слова с цифрами из текста
func extractEnglishWords(text string) []string {
	re := regexp.MustCompile(`[a-zA-Z0-9]+`)
	return re.FindAllString(text, -1)
}

func isNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func getTopWords(words map[string]int, limit int) []*entity.TopWords {
	top := []*entity.TopWords{}

	for word, mentions := range words {
		top = append(top, &entity.TopWords{Word: word, Mentions: mentions})
	}

	sort.Slice(top, func(i, j int) bool {
		return top[i].Mentions > top[j].Mentions
	})

	if len(top) < limit {
		return top
	}

	return top[:limit]
}
