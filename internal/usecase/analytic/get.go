package analyticUC

import (
	"context"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/lib/paging"
	"github.com/himmel520/uoffer/require/internal/usecase"
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
