package positions

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (r *PositionsRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	query, args, err := squirrel.Delete("posts").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)
	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrPostNotFound
	}
	return err

}
