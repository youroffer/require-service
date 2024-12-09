package filterUC

import (
	"context"
)

func (uc *FilterUC) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, uc.db.DB(), id)
}
