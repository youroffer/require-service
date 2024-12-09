package filterUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

type (
	FilterUC struct {
		db   DBXT
		repo FilterRepo
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	FilterRepo interface {
		Create(ctx context.Context, qe repository.Querier, filter string) (*entity.Filter, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
	}
)

func New(db DBXT, repo FilterRepo) *FilterUC {
	return &FilterUC{db: db, repo: repo}
}
