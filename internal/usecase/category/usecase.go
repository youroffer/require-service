package categoryUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

type (
	CategoryUC struct {
		db   DBXT
		repo CategoryRepo
	}

	DBXT interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	CategoryRepo interface {
		Create(ctx context.Context, qe repository.Querier, category string) (*entity.Category, error)
	}
)

func New(db DBXT, repo CategoryRepo) *CategoryUC {
	return &CategoryUC{db: db, repo: repo}
}