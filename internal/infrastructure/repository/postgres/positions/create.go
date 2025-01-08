package positions

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *PositionsRepo) Create(ctx context.Context, qe repository.Querier, post *entity.Position) (*entity.PositionResp, error) {
	query, args, err := squirrel.Insert("posts").
		Columns(
			"categories_id",
			"logo_id",
			"title",
			"public").
		Values(
			post.CategoriesID,
			post.LogoID,
			post.Title,
			post.Public,
		).
		PlaceholderFormat(squirrel.Dollar).
		Suffix(`returning id`).
		ToSql()
	if err != nil {
		return nil, err
	}

	var id int
	err = qe.QueryRow(ctx, query, args...).Scan(&id)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) {
			if pgErr.Code == repoerr.FKViolation {
				return nil, repoerr.ErrPostExists
			}
		}
		return nil, err
	}

	fmt.Println("query", query)

	return r.GetByID(ctx, qe, id)
}
