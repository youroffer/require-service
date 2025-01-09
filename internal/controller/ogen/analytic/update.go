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

func (h *Handler) V1AdminAnalyticsAnalyticIDPut(ctx context.Context, req *api.AnalyticPut, params api.V1AdminAnalyticsAnalyticIDPutParams) (api.V1AdminAnalyticsAnalyticIDPutRes, error) {
	newAnalytic := &entity.AnalyticUpdate{
		PostID:      entity.Optional[int]{Value: req.GetPostsID().Value, Set: req.GetPostsID().IsSet()},
		SearchQuery: entity.Optional[string]{Value: req.GetSearchQuery().Value, Set: req.GetSearchQuery().IsSet()},
	}

	if !newAnalytic.IsSet() {
		return &api.V1AdminAnalyticsAnalyticIDPutBadRequest{Message: "no changes"}, nil
	}

	ad, err := h.uc.Update(ctx, params.AnalyticID, newAnalytic)
	switch {
	case errors.Is(err, repoerr.ErrAnalyticNotFound):
		return &api.V1AdminAnalyticsAnalyticIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrAnalyticExist):
		return &api.V1AdminAnalyticsAnalyticIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return entity.AnalyticRespToApi(ad), nil
}
