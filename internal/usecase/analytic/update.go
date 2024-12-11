package analyticUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
)

func (uc *AnalyticUC) Update(ctx context.Context, id int, analytic *entity.AnalyticUpdate) (*entity.AnalyticResp, error) {
	return uc.repo.Update(ctx, uc.db.DB(), id, analytic)

}
