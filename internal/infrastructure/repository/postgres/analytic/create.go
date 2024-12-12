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
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *AnalyticRepo) Create(ctx context.Context, qe repository.Querier, analytic *entity.Analytic) (*entity.AnalyticResp, error) {
	query, args, err := squirrel.
		Insert("analytics").
		Columns("posts_id", "search_query").
		Values(analytic.PostID, analytic.SearchQuery).
		Suffix("RETURNING id, posts_id, search_query, parse_at, vacancies_num").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var (
		analyticResp = &entity.AnalyticResp{}
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

	analyticResp.ParseAt = entity.Optional[time.Time]{Value: parseAt.Time, Set: parseAt.Valid}
	analyticResp.VacanciesNum = entity.Optional[int]{Value: int(vacanciesNum.Int64), Set: vacanciesNum.Valid}

	return analyticResp, nil
}
