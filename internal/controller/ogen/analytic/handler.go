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
	}
)

func New(uc AnalyticUc) *Handler {
	return &Handler{
		uc: uc,
	}
}
