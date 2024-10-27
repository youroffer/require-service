package usecase

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache/errcache"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"

	"github.com/sirupsen/logrus"
)

type CategoryUsecase struct {
	repo  repository.CategoryRepo
	cache cache.CategoryCache
	log   *logrus.Logger
}

func NewCategoryUsecase(repo repository.CategoryRepo, cache cache.CategoryCache, log *logrus.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, cache: cache, log: log}
}

func (uc *CategoryUsecase) GetAllWithPublicPosts(ctx context.Context) (map[string][]*entity.PostResponse, error) {
	// проверка кэша
	categories, err := uc.cache.GetAllWithPublicPosts(ctx)
	if err != nil {
		if !errors.Is(err, errcache.ErrKeyNotFound) {
			uc.log.Error(err)
		}

		categories, err = uc.repo.GetAllWithPosts(ctx, true)
		if err != nil {
			return nil, err
		}

		// сохраняем в редис, даже при отмене контекста
		err = uc.cache.SetAllWithPublicPosts(context.Background(), categories)
		if err != nil {
			uc.log.Error(err)
		}
	}

	return categories, nil
}

func (uc *CategoryUsecase) GetAllWithPosts(ctx context.Context) (map[string][]*entity.PostResponse, error) {
	return uc.repo.GetAllWithPosts(ctx, false)
}

func (uc *CategoryUsecase) GetAll(ctx context.Context) ([]*entity.Category, error) {
	return uc.repo.GetAll(ctx)
}

func (uc *CategoryUsecase) Add(ctx context.Context, category *entity.Category) (*entity.Category, error) {
	return uc.repo.Add(ctx, category)
}

func (uc *CategoryUsecase) Update(ctx context.Context, category, title string) (*entity.Category, error) {
	return uc.repo.Update(ctx, category, title)
}

func (uc *CategoryUsecase) Delete(ctx context.Context, category string) error {
	return uc.repo.Delete(ctx, category)
}

func (uc *CategoryUsecase) DeleteCache(ctx context.Context) error {
	return uc.cache.DeleteAllWithPublicPosts(ctx)
}
