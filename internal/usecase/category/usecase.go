package categoryUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

type (
	CategoryUC struct {
		db   DBTX
		repo CategoryRepo
	}

	DBTX interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	CategoryRepo interface {
		Create(ctx context.Context, qe repository.Querier, category *entity.Category) (*entity.Category, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Count(ctx context.Context, qe repository.Querier) (int, error)
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Category, error)
		Update(ctx context.Context, qe repository.Querier, id int, category *entity.CategoryUpdate) (*entity.Category, error)
	}
)

func New(db DBTX, repo CategoryRepo) *CategoryUC {
	return &CategoryUC{db: db, repo: repo}
}
