package category

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminCategoriesPost(ctx context.Context, req *api.CategoryPost) (api.V1AdminCategoriesPostRes, error) {
	category, err := h.uc.Create(ctx, req.GetTitle())

	switch {
	case errors.Is(err, repoerr.ErrCategoryExists):
		return &api.V1AdminCategoriesPostConflict{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return entity.ConvertCategoryToApi(category), nil
}
