package category

import (
	"context"
	"errors"

	"github.com/go-chi/chi/middleware"
	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/rs/zerolog/log"
	logSetup "github.com/youroffer/logger"
)

func (h *Handler) V1AdminCategoriesCategoryIDDelete(ctx context.Context, params api.V1AdminCategoriesCategoryIDDeleteParams) (api.V1AdminCategoriesCategoryIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.CategoryID)

	switch {
	case errors.Is(err, repoerr.ErrCategoryNotFound):
		return &api.V1AdminCategoriesCategoryIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		log.Err(err).Str(logSetup.RequestID, middleware.GetReqID(ctx))
		return nil, err
	}

	return &api.V1AdminCategoriesCategoryIDDeleteNoContent{}, nil
}
