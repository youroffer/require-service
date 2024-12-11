package category

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/controller/ogen"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

func (h *Handler) V1AdminCategoriesGet(ctx context.Context, params api.V1AdminCategoriesGetParams) (api.V1AdminCategoriesGetRes, error) {
	CategoriesResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrCategoryNotFound):
		return &api.V1AdminCategoriesGetNotFound{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return CategoriesResp.ToApi(), nil
}
