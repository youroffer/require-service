package CategoryRepo

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (r *CategoryRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.Category, error) {
	builder := squirrel.
		Select(
			"id", 
			"title", 
			"public").
		From("categories").
		OrderBy("title").
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

	categories := []*entity.Category{}
	for rows.Next() {
		category := &entity.Category{}
		if err := rows.Scan(&category.ID, &category.Title, &category.Public); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if len(categories) == 0 {
		return nil, repoerr.ErrCategoryNotFound
	}

	return categories, nil
}
