package analyticUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

type (
	AnalyticUC struct {
		db   DBTX
		repo AnalyticRepo
	}

	DBTX interface {
		DB() repository.Querier
		InTransaction(ctx context.Context, fn repository.TransactionFunc) error
	}

	AnalyticRepo interface {
		Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.AnalyticResp, error)
		Count(ctx context.Context, qe repository.Querier) (int, error)
		Create(ctx context.Context, qe repository.Querier, analytic *entity.Analytic) (*entity.AnalyticResp, error)
		Delete(ctx context.Context, qe repository.Querier, id int) error
		Update(ctx context.Context, qe repository.Querier, id int, analytic *entity.AnalyticUpdate) (*entity.AnalyticResp, error)
	}
)

func New(db DBTX, repo AnalyticRepo) *AnalyticUC {
	return &AnalyticUC{db: db, repo: repo}
}
