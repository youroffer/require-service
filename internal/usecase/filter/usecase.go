package filterUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

type (
	FilterUC struct {
		db   DBTX
		repo FilterRepo
	}

	DBTX interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	FilterRepo interface {
		Create(ctx context.Context, qe repository.Querier, filter string) (*entity.Filter, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Filter, error)
		Count(ctx context.Context, qe repository.Querier) (int, error)
	}
)

func New(db DBTX, repo FilterRepo) *FilterUC {
	return &FilterUC{db: db, repo: repo}
}
