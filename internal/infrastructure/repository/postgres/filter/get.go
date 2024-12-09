package filterRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (r *FilterRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Filter, error) {
	builder := squirrel.
		Select(
			"id",
			"word").
		From("filters").
		OrderBy("word").
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %w", err)
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	filters := []*entity.Filter{}
	for rows.Next() {
		filter := &entity.Filter{}
		if err := rows.Scan(&filter.ID, &filter.Word); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}

	if len(filters) == 0 {
		return nil, repoerr.ErrFilterNotFound
	}

	return filters, nil
}
