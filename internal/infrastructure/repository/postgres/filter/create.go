package filterRepo

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *FilterRepo) Create(ctx context.Context, qe repository.Querier, filter string) (*entity.Filter, error) {
	newFilter := &entity.Filter{}

	err := qe.QueryRow(ctx, `
	insert into filters 
		(word) 
	values 
		($1) 
	returning *;`, filter).Scan(
		&newFilter.ID, &newFilter.Word)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.UniqueConstraint {
			return nil, repoerr.ErrFilterExist
		}
	}

	return newFilter, err
}
