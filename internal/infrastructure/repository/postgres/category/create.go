package CategoryRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *CategoryRepo) Create(ctx context.Context, qe repository.Querier, category string) (*entity.Category, error) {
	query, args, err := squirrel.
		Insert("categories").
		Columns("title", "public").
		Values(category, false).
		Suffix("returning id, title, public").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	newCategory := &entity.Category{}

	err = qe.QueryRow(ctx, query, args...).Scan(
		&newCategory.ID,
		&newCategory.Title,
		&newCategory.Public)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.UniqueConstraint {
			return nil, repoerr.ErrCategoryExists
		}
	}

	return newCategory, err
}
