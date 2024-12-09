package filter

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/controller/ogen"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

func (h *Handler) V1AdminFiltersGet(ctx context.Context, params api.V1AdminFiltersGetParams) (api.V1AdminFiltersGetRes, error) {
	filtersResp, err := h.uc.Get(ctx, usecase.PageParams{
		Page:    uint64(params.Page.Or(ogen.Page)),
		PerPage: uint64(params.PerPage.Or(ogen.PerPage)),
	})

	switch {
	case errors.Is(err, repoerr.ErrFilterNotFound):
		return &api.V1AdminFiltersGetNotFound{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}

	return filtersResp.ToApi(), nil
}
