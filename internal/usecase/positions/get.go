package positions

import (
	"context"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/lib/paging"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

func (uc *PositionUC) Get(ctx context.Context, params usecase.PageParams) (*entity.PositionsResp, error) {
	position, err := uc.repo.Get(ctx, uc.db.DB(), repository.PaginationParams{
		Limit:  entity.Optional[uint64]{Value: params.PerPage, Set: true},
		Offset: entity.Optional[uint64]{Value: params.Page * params.PerPage, Set: true},
	})
	if err != nil {
		return nil, fmt.Errorf("repo get: %w", err)
	}

	count, err := uc.repo.Count(ctx, uc.db.DB())
	if err != nil {
		return nil, fmt.Errorf("repo count: %w", err)
	}

	return &entity.PositionsResp{
		Data:    position,
		Page:    params.Page,
		Pages:   paging.CalculatePages(count, params.PerPage),
		PerPage: params.PerPage,
	}, err
}
