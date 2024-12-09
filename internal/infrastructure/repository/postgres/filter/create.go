package filterRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *FilterRepo) Create(ctx context.Context, qe repository.Querier, filter string) (*entity.Filter, error) {
	query, args, err := squirrel.
		Insert("filters").
		Columns("word").
		Values(filter).
		Suffix("returning id, word").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	newFilter := &entity.Filter{}
	err = qe.QueryRow(ctx, query, args...).Scan(
		&newFilter.ID,
		&newFilter.Word)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
		return nil, repoerr.ErrFilterExist
	}

	return newFilter, err
}
