package analytic

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/controller/ogen"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

func (h *Handler) V1AdminAnalyticsGet(ctx context.Context, params api.V1AdminAnalyticsGetParams) (api.V1AdminAnalyticsGetRes, error) {
	analyticsResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrAnalyticNotFound):
		return &api.V1AdminAnalyticsGetNotFound{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return analyticsResp.ToApi(), nil
}
