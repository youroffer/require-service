package filter

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

type (
	Handler struct {
		uc FilterUc
	}

	FilterUc interface {
		Create(ctx context.Context, filter string) (*entity.Filter, error)
		Delete(ctx context.Context, id int) error
		Get(ctx context.Context, params usecase.PageParams) (*entity.FiltersResp, error)
	}
)

func New(uc FilterUc) *Handler {
	return &Handler{
		uc: uc,
	}
}
