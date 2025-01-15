package analytic

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/controller/ogen"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/uoffer/require/internal/usecase"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
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
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return analyticsResp.ToApi(), nil
}

func (h *Handler) V1AnalyticsAnalyticIDGet(ctx context.Context, params api.V1AnalyticsAnalyticIDGetParams) (api.V1AnalyticsAnalyticIDGetRes, error) {
	analytic, err := h.uc.GetWithWordsByID(ctx, params.AnalyticID, false)

	switch {
	case errors.Is(err, repoerr.ErrAnalyticNotFound):
		return &api.V1AnalyticsAnalyticIDGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return analytic.ToApiWithWords(), nil
}

func (h *Handler) V1AnalyticsAnalyticIDLimitGet(ctx context.Context, params api.V1AnalyticsAnalyticIDLimitGetParams) (api.V1AnalyticsAnalyticIDLimitGetRes, error) {
	analytic, err := h.uc.GetWithWordsByID(ctx, params.AnalyticID, true)

	switch {
	case errors.Is(err, repoerr.ErrAnalyticNotFound):
		return &api.V1AnalyticsAnalyticIDLimitGetNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return analytic.ToApiWithWords(), nil
}
