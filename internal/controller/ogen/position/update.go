package position

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

func (h *Handler) V1AdminPositionsPositionIDPut(ctx context.Context, req *api.PositionPut, params api.V1AdminPositionsPositionIDPutParams) (api.V1AdminPositionsPositionIDPutRes, error) {
	newPosition := &entity.PositionUpdate{
		CategoriesID: entity.Optional[int]{Value: req.GetCategoriesID().Value, Set: req.GetCategoriesID().Set},
		LogoID:       entity.Optional[int]{Value: req.GetLogoID().Value, Set: req.GetLogoID().Set},
		Title:        entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Public:       entity.Optional[bool]{Value: req.GetPublic().Value, Set: req.GetPublic().Set},
	}

	if !newPosition.IsSet() {
		return &api.V1AdminPositionsPositionIDPutBadRequest{Message: "no changes"}, nil
	}

	post, err := h.uc.Update(ctx, params.PositionID, newPosition)
	switch {
	case errors.Is(err, repoerr.ErrPostNotFound):
		return &api.V1AdminPositionsPositionIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrPostExists):
		return &api.V1AdminPositionsPositionIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return entity.PositionRespToApi(post), nil
}
