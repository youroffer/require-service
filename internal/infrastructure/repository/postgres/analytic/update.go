package analyticRepo

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

func (r *AnalyticRepo) Update(ctx context.Context, qe repository.Querier, id int, analytic *entity.AnalyticUpdate) (*entity.AnalyticResp, error) {
	builder := squirrel.Update("analytics").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	if analytic.PostID.Set {
		builder = builder.Set("posts_id", analytic.PostID.Value)
	}

	if analytic.SearchQuery.Set {
		builder = builder.Set("search_query", analytic.SearchQuery.Value)
	}

	if analytic.VacanciesNum.Set {
		builder = builder.Set("vacancies_num", analytic.VacanciesNum.Value)
	}

	if analytic.ParseAt.Set {
		builder = builder.Set("parse_at", analytic.ParseAt.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	cmdTag, err := qe.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint {
			return nil, repoerr.ErrAnalyticExist
		}

		return nil, fmt.Errorf("get analytics: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, repoerr.ErrAnalyticNotFound
	}

	return r.GetByID(ctx, qe, id)
}
