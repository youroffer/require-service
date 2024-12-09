package filterUC

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
)

func (uc *FilterUC) Create(ctx context.Context, filter string) (*entity.Filter, error) {
	return uc.repo.Create(ctx, uc.db.DB(), filter)
}
