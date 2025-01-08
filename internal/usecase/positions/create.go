package positions

import (
	"context"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
)

func (uc *PositionUC) Create(ctx context.Context, post *entity.Position) (*entity.PositionResp, error) {
	newPosition, err := uc.repo.Create(ctx, uc.db.DB(), post)
	if err != nil {
		return nil, fmt.Errorf("create position: %w", err)
	}

	return newPosition, nil
}
