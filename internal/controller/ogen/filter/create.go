package filter

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

func (h *Handler) V1AdminFiltersPost(ctx context.Context, req *api.V1AdminFiltersPostReq) (api.V1AdminFiltersPostRes, error) {
	filter, err := h.uc.Create(ctx, req.GetWord())

	switch {
	case errors.Is(err, repoerr.ErrFilterExist):
		return &api.V1AdminFiltersPostConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return entity.ConvertFilterToApi(filter), nil
}
