package positions

import (
	"context"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
)

func (uc *PositionUC) Update(ctx context.Context, id int, post *entity.PositionUpdate) (*entity.PositionResp, error) {
	updatedPosition, err := uc.repo.Update(ctx, uc.db.DB(), id, post)
	if err != nil {
		return nil, fmt.Errorf("update adv: %w", err)
	}

	return updatedPosition, nil
}
