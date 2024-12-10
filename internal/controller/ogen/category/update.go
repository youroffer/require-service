package category

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	log "github.com/youroffer/logger"
)

func (h *Handler) V1AdminCategoriesCategoryIDPut(ctx context.Context, req *api.CategoryPut, params api.V1AdminCategoriesCategoryIDPutParams) (api.V1AdminCategoriesCategoryIDPutRes, error) {

	newCategory := &entity.CategoryUpdate{
		Title:  entity.Optional[string]{Value: req.GetTitle().Value, Set: req.GetTitle().Set},
		Public: entity.Optional[bool]{Value: req.GetPublic().Value, Set: req.GetPublic().Set},
	}

	if !newCategory.IsSet() {
		return &api.V1AdminCategoriesCategoryIDPutBadRequest{Message: "no changes"}, nil
	}

	category, err := h.uc.Update(ctx, params.CategoryID, newCategory)
	switch {
	case errors.Is(err, repoerr.ErrCategoryNotFound):
		return &api.V1AdminCategoriesCategoryIDPutNotFound{Message: err.Error()}, nil
	case errors.Is(err, repoerr.ErrCategoryExists):
		return &api.V1AdminCategoriesCategoryIDPutConflict{Message: err.Error()}, nil
	case err != nil:
		log.ErrFields(err, log.Fields{
			log.RequestID: middleware.GetReqID(ctx),
		})
		return nil, err
	}

	return entity.ConvertCategoryToApi(category), nil
}
