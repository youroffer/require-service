package analytic

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
)

func (h *Handler) V1AdminAnalyticsPost(ctx context.Context, req *api.AnalyticPost) (api.V1AdminAnalyticsPostRes, error) {
	analytic, err := h.uc.Create(ctx, &entity.Analytic{
		PostID:      req.GetPostID(),
		SearchQuery: req.GetSearchQuery(),
	})

	switch {
	case errors.Is(err, repoerr.ErrAnalyticDependencyNotFound):
		return &api.V1AdminAnalyticsPostConflict{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrAnalyticExist):
		return &api.V1AdminAnalyticsPostUnprocessableEntity{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return entity.AnalyticRespToApi(analytic), nil
}
