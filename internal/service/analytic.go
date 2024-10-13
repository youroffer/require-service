package service

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
	"github.com/himmel520/uoffer/require/roles"
)

func (s *Service) clearAnalyticCache(ctx context.Context, postID int) {
	if err := s.cache.DeleteAnalyticWithWords(ctx, postID); err != nil {
		s.log.Error(err)
	}
}

func (s *Service) AddAnalytic(ctx context.Context, analytic *models.Analytic) (*models.Analytic, error) {
	newAnalytic, err := s.repo.AddAnalytic(ctx, analytic)
	if err != nil {
		return nil, err
	}

	go s.clearAnalyticCache(context.Background(), newAnalytic.PostID)

	return newAnalytic, nil
}

func (s *Service) UpdateAnalytic(ctx context.Context, id int, analytic *models.AnalyticUpdate) (*models.Analytic, error) {
	newAnalytic, err := s.repo.UpdateAnalytic(ctx, id, analytic)
	if err != nil {
		return nil, err
	}

	go s.clearAnalyticCache(context.Background(), newAnalytic.PostID)

	return newAnalytic, nil
}

func (s *Service) DeleteAnalytic(ctx context.Context, id int) error {
	postID, err := s.repo.GetPostIDByAnalytic(ctx, id)
	if err != nil {
		return err
	}

	err = s.repo.DeleteAnalytic(ctx, id)
	if err != nil {
		return err
	}

	go s.clearAnalyticCache(context.Background(), postID)

	return nil
}

func (s *Service) GetAnalyticWithWords(ctx context.Context, postID int, role string) (*models.AnalyticWithWords, error) {
	var limit bool
	if role == roles.Anonym {
		limit = true
	}

	analytic, err := s.cache.GetAnalyticWithWords(ctx, postID)
	if err != nil {
		if !errors.Is(err, repository.ErrAnalyticNotFound) {
			s.log.Error(err)
		}

		analytic, err := s.repo.GetAnalytic(ctx, postID)
		if err != nil {
			return nil, err
		}

		analyticWords := &models.AnalyticWithWords{
			Analytic: analytic,
			Skills:   []*models.TopWords{},
			Keywords: []*models.TopWords{},
		}

		go func() {
			if err := s.cache.SetAnalyticWithWords(context.Background(), analyticWords, postID); err != nil {
				s.log.Error(err)
			}
		}()

		return analyticWords, nil
	}

	// проверка len во избежания паники(а вдруг)
	if limit && len(analytic.Keywords) > 8 && len(analytic.Keywords) > 8 {
		// ограничиваем количество слов для анонимов
		analytic.Keywords = analytic.Keywords[:8]
		analytic.Skills = analytic.Skills[:8]
	}

	return analytic, nil
}
