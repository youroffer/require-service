package analyticUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
)

func (uc *AnalyticUC) Create(ctx context.Context, analytic *entity.Analytic) (*entity.AnalyticResp, error) {
	return uc.repo.Create(ctx, uc.db.DB(), analytic)
}
