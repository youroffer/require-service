package position

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
)

func (h *Handler) V1AdminPositionsPositionIDDelete(ctx context.Context, params api.V1AdminPositionsPositionIDDeleteParams) (api.V1AdminPositionsPositionIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.PositionID)

	switch {
	case errors.Is(err, repoerr.ErrPostNotFound):
		return &api.V1AdminPositionsPositionIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return &api.V1AdminPositionsPositionIDDeleteOK{}, nil
}
