package filter

import (
	"context"

	api "github.com/himmel520/uoffer/require/api/oas"
)

func (h *Handler) V1AdminFiltersGet(ctx context.Context, params api.V1AdminFiltersGetParams) (api.V1AdminFiltersGetRes, error) {
	return &api.FiltersResp{}, nil
}
