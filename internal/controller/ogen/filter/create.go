package filter

import (
	"context"
	"errors"

	api "github.com/himmel520/uoffer/require/api/oas"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (h *Handler) V1AdminFiltersPost(ctx context.Context, req *api.V1AdminFiltersPostReq) (api.V1AdminFiltersPostRes, error) {
	filter, err := h.uc.Create(ctx, req.GetWord())
	
	switch {
	case errors.Is(err, repoerr.ErrFilterExist):
		return &api.V1AdminFiltersPostConflict{Message: err.Error()}, nil
	case err != nil:
		return nil, err
	}
	
	return entity.ConvertFilterToApi(filter), nil
}
