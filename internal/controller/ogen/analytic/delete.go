package analytic

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
)

func (h *Handler) V1AdminAnalyticsAnalyticIDDelete(ctx context.Context, params api.V1AdminAnalyticsAnalyticIDDeleteParams) (api.V1AdminAnalyticsAnalyticIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.AnalyticID)

	switch {
	case errors.Is(err, repoerr.ErrAnalyticNotFound):
		return &api.V1AdminAnalyticsAnalyticIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return &api.V1AdminAnalyticsAnalyticIDDeleteNoContent{}, nil
}
