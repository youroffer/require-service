package category

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
)

type (
	Handler struct {
		uc CategoryUC
	}

	CategoryUC interface {
		Create(ctx context.Context, category string) (*entity.Category, error)
	}
)

func New(uc CategoryUC) *Handler {
	return &Handler{
		uc: uc,
	}
}
