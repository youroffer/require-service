package filter

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminFiltersFilterIDDelete(ctx context.Context, params api.V1AdminFiltersFilterIDDeleteParams) (api.V1AdminFiltersFilterIDDeleteRes, error) {
	err := h.uc.Delete(ctx, params.FilterID)
	switch {
	case errors.Is(err, repoerr.ErrFilterNotFound):
		return &api.V1AdminFiltersFilterIDDeleteNotFound{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}
	return &api.V1AdminFiltersFilterIDDeleteNoContent{}, nil
}
