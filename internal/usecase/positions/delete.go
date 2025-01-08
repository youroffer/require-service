package positions

import (
	"context"
	"fmt"
)

func (uc *PositionUC) Delete(ctx context.Context, id int) error {
	if err := uc.repo.Delete(ctx, uc.db.DB(), id); err != nil {
		return fmt.Errorf("delete position: %w", err)
	}

	return nil
}
