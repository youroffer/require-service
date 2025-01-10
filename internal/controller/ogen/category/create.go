package category

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

func (h *Handler) V1AdminCategoriesPost(ctx context.Context, req *api.CategoryPost) (api.V1AdminCategoriesPostRes, error) {
	category, err := h.uc.Create(ctx, &entity.Category{
		Title:  req.GetTitle(),
		Public: req.GetPublic()})

	switch {
	case errors.Is(err, repoerr.ErrCategoryExists):
		return &api.V1AdminCategoriesPostConflict{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return entity.ConvertCategoryToApi(category), nil
}
