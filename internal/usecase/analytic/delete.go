package analyticUC

import "context"

func (uc *AnalyticUC) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, uc.db.DB(), id)
}
