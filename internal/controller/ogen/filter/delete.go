package filter

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
)

func (h *Handler) V1AdminFiltersFilterIDDelete(ctx context.Context, params api.V1AdminFiltersFilterIDDeleteParams) (api.V1AdminFiltersFilterIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.FilterID)

	switch {
	case errors.Is(err, repoerr.ErrFilterNotFound):
		return &api.V1AdminFiltersFilterIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return &api.V1AdminFiltersFilterIDDeleteNoContent{}, nil
}
