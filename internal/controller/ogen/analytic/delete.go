package analytic

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminAnalyticsAnalyticIDDelete(ctx context.Context, params api.V1AdminAnalyticsAnalyticIDDeleteParams) (api.V1AdminAnalyticsAnalyticIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.AnalyticID)

	switch {
	case errors.Is(err, repoerr.ErrAnalyticNotFound):
		return &api.V1AdminAnalyticsAnalyticIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return &api.V1AdminAnalyticsAnalyticIDDeleteNoContent{}, nil
}
