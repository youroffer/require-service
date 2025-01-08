package positions

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

type (
	PositionUC struct {
		db   DBXT
		repo PositionRepo
	}

	DBXT interface {
		DB() repository.Querier
	}

	PositionRepo interface {
		Create(ctx context.Context, qe repository.Querier, post *entity.Position) (*entity.PositionResp, error)
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.PositionResp, error)
		Count(ctx context.Context, qe repository.Querier) (int, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
	}
)

func New(db DBXT, repo PositionRepo) *PositionUC {
	return &PositionUC{
		db:   db,
		repo: repo,
	}
}
