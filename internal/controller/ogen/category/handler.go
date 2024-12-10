package category

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/usecase"
)

type (
	Handler struct {
		uc CategoryUC
	}

	CategoryUC interface {
		Create(ctx context.Context, category string) (*entity.Category, error)
		Delete(ctx context.Context, id int) error
		Get(ctx context.Context, params usecase.PageParams) (*entity.CategoriesResp, error)
		Update(ctx context.Context, id int, category *entity.CategoryUpdate) (*entity.Category, error)
	}
)

func New(uc CategoryUC) *Handler {
	return &Handler{
		uc: uc,
	}
}
