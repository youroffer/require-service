package analyticUC

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/lib/paging"
	"github.com/himmel520/uoffer/require/internal/usecase"
	"github.com/rs/zerolog/log"
)

func (uc *AnalyticUC) Get(ctx context.Context, params usecase.PageParams) (*entity.AnalyticsResp, error) {
	analytics, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  entity.NewOptional(params.PerPage),
		Offset: entity.NewOptional(params.Page * params.PerPage)})
	if err != nil {
		return nil, fmt.Errorf("repo get: %w", err)
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, fmt.Errorf("repo count: %w", err)
	}

	return &entity.AnalyticsResp{
		Data:    analytics,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

func (uc *AnalyticUC) GetWithWordsByID(ctx context.Context, analyticID int, limit bool) (*entity.AnalyticWithWords, error) {
	cacheData, err := uc.cache.Get(ctx, fmt.Sprintf(cache.AnalyticKeyFmt, analyticID))
	if err != nil {
		if !errors.Is(err, cache.ErrKeyNotFound) {
			log.Err(err)
		}

		analytic, err := uc.repo.GetByID(ctx, uc.db.DB(), analyticID)
		if err != nil {
			return nil, fmt.Errorf("repo get by ID: %w", err)
		}

		return &entity.AnalyticWithWords{
			Analytic: analytic,
			Skills:   []*entity.TopWords{},
			Keywords: []*entity.TopWords{}}, nil
	}

	analytic := &entity.AnalyticWithWords{}
	if err = json.Unmarshal([]byte(cacheData), analytic); err != nil {
		return nil, fmt.Errorf("funmarshal cache: %w", err)
	}

	if limit && len(analytic.Keywords) > 8 && len(analytic.Keywords) > 8 {
		analytic.Keywords = analytic.Keywords[:8]
		analytic.Skills = analytic.Skills[:8]
	}

	return analytic, err
}
