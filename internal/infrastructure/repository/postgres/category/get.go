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

func (r *CategoryRepo) GetPublic(ctx context.Context, qe repository.Querier) (entity.CategoriesPublicPostsResp, error) {
	query, args, err := squirrel.
		Select(
			"c.title",
			"p.id",
			"p.logo_id",
			"p.title",
			"p.public").
		From("categories AS c").
		Join("posts AS p ON p.categories_id = c.id").
		OrderBy("c.title").
		Where(squirrel.Eq{"c.public": true}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := entity.CategoriesPublicPostsResp{}
	for rows.Next() {
		var (
			category string
			position entity.CategoryPosition
		)

		if err := rows.Scan(
			&category,
			&position.ID,
			&position.LogoID,
			&position.Title,
			&position.Public); err != nil {
			return nil, err
		}

		categories[category] = append(categories[category], position)
	}

	if len(categories) == 0 {
		return nil, repoerr.ErrCategoryNotFound
	}

	return categories, err
}
