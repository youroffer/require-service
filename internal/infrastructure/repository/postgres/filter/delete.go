package filterRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (r *FilterRepo) Delete(ctx context.Context, qe repository.Querier, id int) error {
	query, args, err := squirrel.Delete("filters").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	fmt.Println(query, "\n", args)

	cmdTag, err := qe.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("row exec: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrFilterNotFound
	}

	return err
}