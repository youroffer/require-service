package position

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/rs/zerolog/log"

	logSetup "github.com/youroffer/logger"
)

func (h *Handler) V1AdminPositionsPost(ctx context.Context, req *api.PositionPost) (api.V1AdminPositionsPostRes, error) {
	post, err := h.uc.Create(ctx, &entity.Position{
		CategoriesID: req.GetCategoriesID(),
		LogoID:       req.GetLogoID(),
		Title:        req.GetTitle(),
		Public:       req.GetPublic(),
	})

	switch {
	case errors.Is(err, repoerr.ErrPostDependencyNotFound):
		return &api.V1AdminPositionsPostConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	fmt.Println("handler", post)

	return entity.PositionRespToApi(post), nil
}
