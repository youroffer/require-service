package analyticRepo

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AnalyticRepo) Update(ctx context.Context, qe repository.Querier, id int, analytic *entity.AnalyticUpdate) (*entity.AnalyticResp, error) {
	builder := squirrel.Update("analytics").
		Where(squirrel.Eq{"id": id}).
		Suffix(`RETURNING id, posts_id, search_query, parse_at, vacancies_num`).
		PlaceholderFormat(squirrel.Dollar)

	if analytic.PostID.Set {
		builder = builder.Set("posts_id", analytic.PostID.Value)
	}

	if analytic.SearchQuery.Set {
		builder = builder.Set("search_query", analytic.SearchQuery.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var (
		analyticResp entity.AnalyticResp
		parseAt      sql.NullTime
		vacanciesNum sql.NullInt64
	)

	err = qe.QueryRow(ctx, query, args...).Scan(
		&analyticResp.ID,
		&analyticResp.PostTitle,
		&analyticResp.SearchQuery,
		&parseAt,
		&vacanciesNum,
	)
	if err != nil {
		var pgErr *pgconn.PgError
		switch {
		case errors.Is(err, pgx.ErrNoRows):
			return nil, repoerr.ErrAnalyticNotFound
		case errors.As(err, &pgErr) && pgErr.Code == repoerr.UniqueConstraint:
			return nil, repoerr.ErrAnalyticExist
		}
	}

	analyticResp.ParseAt = entity.Optional[time.Time]{Value: parseAt.Time, Set: parseAt.Valid}
	analyticResp.VacanciesNum = entity.Optional[int]{Value: int(vacanciesNum.Int64), Set: vacanciesNum.Valid}

	return &analyticResp, nil
}
