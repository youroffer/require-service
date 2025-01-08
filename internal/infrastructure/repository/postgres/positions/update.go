package positions

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *PositionsRepo) Update(ctx context.Context, qe repository.Querier, id int, post *entity.PositionUpdate) (*entity.PositionResp, error) {
	builder := squirrel.Update("posts").
		Where(squirrel.Eq{"id": id}).
		Suffix(`returning id`).
		PlaceholderFormat(squirrel.Dollar)

	if post.CategoriesID.Set {
		builder = builder.Set("categories_id", post.CategoriesID.Value)
	}

	if post.LogoID.Set {
		builder = builder.Set("logo_id", post.LogoID.Value)
	}

	if post.Title.Set {
		builder = builder.Set("title", post.Title.Value)
	}

	if post.Public.Set {
		builder = builder.Set("public", post.Public.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)

	var pgErr *pgconn.PgError
	if err != nil {
		if errors.As(err, &pgErr) && pgErr.Code == repoerr.FKViolation {
			return nil, repoerr.ErrPostExists
		}
		return nil, err
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, repoerr.ErrPostNotFound
	}

	return r.GetByID(ctx, qe, id)
}
