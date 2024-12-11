package analytic

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"

	log "github.com/youroffer/logger"
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
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.AnalyticRespToApi(analytic), nil
}
