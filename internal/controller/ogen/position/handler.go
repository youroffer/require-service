package position

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
)

type (
	Handler struct {
		uc PositionUC
	}

	PositionUC interface {
		Create(ctx context.Context, post *entity.Position) (*entity.PositionResp, error)
	}
)

func New(uc PositionUC) *Handler {
	return &Handler{
		uc: uc,
	}
}
