package filterRepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
)

func (r *FilterRepo) Count(ctx context.Context, qe repository.Querier) (int, error) {
	var count int

	query, args, err := squirrel.
		Select("COUNT(*)").
		From("filters").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, err
	}

	err = qe.QueryRow(ctx, query, args...).Scan(&count)
	return count, err
}