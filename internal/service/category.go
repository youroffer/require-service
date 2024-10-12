package service

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
)

func (s *Service) GetCategoriesWithPublicPosts(ctx context.Context) (map[string][]*models.PostResponse, error) {
	// проверка кэша
	categories, err := s.cache.GetCategoriesWithPublicPosts(ctx)
	if err != nil {
		if !errors.Is(err, repository.ErrKeyNotFound) {
			s.log.Error(err)
		}

		categories, err = s.repo.GetCategoriesWithPosts(ctx, true)
		if err != nil {
			return nil, err
		}

		// сохраняем в редис, даже при отмене контекста
		err = s.cache.SetCategoriesWithPublicPosts(context.Background(), categories)
		if err != nil {
			s.log.Error(err)
		}
	}

	return categories, nil
}

func (s *Service) GetCategoriesWithPosts(ctx context.Context) (map[string][]*models.PostResponse, error) {
	return s.repo.GetCategoriesWithPosts(ctx, false)
}

func (s *Service) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	return s.repo.GetAllCategories(ctx)
}

func (s *Service) AddCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	return s.repo.AddCategory(ctx, category)
}

func (s *Service) UpdateCategory(ctx context.Context, category, title string) (*models.Category, error) {
	return s.repo.UpdateCategory(ctx, category, title)
}

func (s *Service) DeleteCategory(ctx context.Context, category string) error {
	return s.repo.DeleteCategory(ctx, category)
}

func (s *Service) DeleteCacheCategoriesAndPosts(ctx context.Context) error {
	return s.cache.DeleteCategoriesWithPublicPosts(ctx)
}
