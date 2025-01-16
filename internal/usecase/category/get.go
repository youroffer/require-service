package categoryUC

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

func (uc *CategoryUC) Get(ctx context.Context, params usecase.PageParams) (*entity.CategoriesResp, error) {
	categories, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  entity.NewOptional(params.PerPage),
		Offset: entity.NewOptional(params.Page * params.PerPage)})
	if err != nil {
		return nil, fmt.Errorf("repo get: %w", err)
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, fmt.Errorf("repo count: %w", err)
	}

	return &entity.CategoriesResp{
		Data:    categories,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}

func (uc *CategoryUC) GetPublic(ctx context.Context) (entity.CategoriesPublicPostsResp, error) {
	var categories entity.CategoriesPublicPostsResp

	bytes, err := uc.cache.Get(ctx, cache.CategoriesWithPostsKey)
	if err != nil {
		if !errors.Is(err, cache.ErrKeyNotFound) {
			log.Err(err)
		}

		categories, err := uc.repo.GetPublic(ctx, uc.db.DB())
		if err != nil {
			return nil, err
		}

		if err = uc.cache.Set(ctx, cache.CategoriesWithPostsKey, categories); err != nil {
			log.Err(err).Msg("set all categories cache")
		}

		return categories, nil
	}

	err = json.Unmarshal([]byte(bytes), &categories)
	return categories, err
}
