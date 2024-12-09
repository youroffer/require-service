package filter

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
)

type (
	Handler struct {
		uc FilterUc
	}

	FilterUc interface {
		Create(ctx context.Context, filter string) (*entity.Filter, error)
	}
)

func New(uc FilterUc) *Handler {
	return &Handler{
		uc: uc,
	}
}
