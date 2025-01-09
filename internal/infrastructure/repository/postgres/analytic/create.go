package analyticRepo

import (
	"context"
	"errors"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AnalyticRepo) Create(ctx context.Context, qe repository.Querier, analytic *entity.Analytic) (*entity.AnalyticResp, error) {
	query, args, err := squirrel.
		Insert("analytics").
		Columns("posts_id", "search_query").
		Values(analytic.PostID, analytic.SearchQuery).
		Suffix(`returning id`).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var analyticID int
	err = qe.QueryRow(ctx, query, args...).Scan(&analyticID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			switch {
			case pgErr.Code == repoerr.FKViolation:
				return nil, repoerr.ErrAnalyticDependencyNotFound
			case pgErr.Code == repoerr.UniqueConstraint:
				return nil, repoerr.ErrAnalyticExist
			}
		}
		return nil, err
	}

	return r.GetByID(ctx, qe, analyticID)
}
