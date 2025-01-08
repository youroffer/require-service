package position

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

type (
	Handler struct {
		uc PositionUC
	}

	PositionUC interface {
		Create(ctx context.Context, post *entity.Position) (*entity.PositionResp, error)
		Get(ctx context.Context, params usecase.PageParams) (*entity.PositionsResp, error)
		Delete(ctx context.Context, id int) error
		Update(ctx context.Context, id int, post *entity.PositionUpdate) (*entity.PositionResp, error)
	}
)

func New(uc PositionUC) *Handler {
	return &Handler{
		uc: uc,
	}
}
