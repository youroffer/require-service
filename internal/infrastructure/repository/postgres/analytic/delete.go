package analyticRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (r *AnalyticRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	query, args, err := squirrel.Delete("analytics").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("row exec: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrAnalyticNotFound
	}

	return err
}
