package categoryUC

import "context"

func (uc *CategoryUC) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, uc.db.DB(), id)
}
