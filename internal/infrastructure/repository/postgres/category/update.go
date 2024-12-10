package CategoryRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CategoryRepo) Update(ctx context.Context, qe repository.Querier, id int, category *entity.CategoryUpdate) (*entity.Category, error) {
	builder := squirrel.
		Update("categories").
		Where(squirrel.Eq{"id": id}).
		Suffix(`
	returning 
	id, 
	title, 
	public`).
		PlaceholderFormat(squirrel.Dollar)

	if category.Title.Set {
		builder = builder.Set("title", category.Title.Value)
	}

	if category.Public.Set {
		builder = builder.Set("public", category.Public.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	newCategory := &entity.Category{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newCategory.ID,
		&newCategory.Title,
		&newCategory.Public)

	var pgErr *pgconn.PgError
	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return nil, repoerr.ErrCategoryNotFound
	case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
		return nil, repoerr.ErrCategoryExists
	}

	return newCategory, err
}
