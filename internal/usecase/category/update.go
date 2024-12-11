package categoryUC

import (
	"context"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
)

func (uc *CategoryUC) Update(ctx context.Context, id int, category *entity.CategoryUpdate) (*entity.Category, error) {
	updatedCategory, err := uc.repo.Update(ctx, uc.db.DB(), id, category)
	if err != nil {
		return nil, fmt.Errorf("update category: %w", err)
	}

	return updatedCategory, nil
}
