package category

import (
	"context"
	"errors"
	"fmt"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/controller/ogen"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/uoffer/require/internal/usecase"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
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
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return CategoriesResp.ToApi(), nil
}

func (h *Handler) V1CategoriesGet(ctx context.Context) (api.V1CategoriesGetRes, error) {
	categories, err := h.uc.GetPublic(ctx)
	if err != nil {
		return nil, fmt.Errorf("categories get: %w", err)
	}

	return entity.CategoryPublicToApi(categories), nil
}
