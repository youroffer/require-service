package analytic

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

type (
	Handler struct {
		uc AnalyticUc
	}

	AnalyticUc interface {
		Get(ctx context.Context, params usecase.PageParams) (*entity.AnalyticsResp, error)
		Create(ctx context.Context, analytic *entity.Analytic) (*entity.AnalyticResp, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, analytic *entity.AnalyticUpdate) (*entity.AnalyticResp, error)
		GetWithWordsByID(ctx context.Context, analyticID int, limit bool) (*entity.AnalyticWithWords, error)
	}
)

func New(uc AnalyticUc) *Handler {
	return &Handler{
		uc: uc,
	}
}
