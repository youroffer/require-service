package analyticRepo

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
)

func (r *AnalyticRepo) Get(ctx context.Context, qe repository.Querier, params repository.PaginationParams) ([]*entity.AnalyticResp, error) {
	builder := squirrel.Select(
		"a.id",
		"p.id",
		"a.search_query",
		"a.parse_at",
		"a.vacancies_num").
		From("analytics AS a").
		Join("posts AS p ON p.id = a.posts_id").
		OrderBy("a.id").
		Limit(params.Limit.Value).
		Offset(params.Offset.Value).
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := qe.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	analyticResp := []*entity.AnalyticResp{}
	for rows.Next() {
		var (
			parseAt      sql.NullTime
			vacanciesNum sql.NullInt64
		)
		
		analytic := &entity.AnalyticResp{}
		if err := rows.Scan(
			&analytic.ID,
			&analytic.PostTitle,
			&analytic.SearchQuery,
			&parseAt,
			&vacanciesNum); err != nil {
			return nil, err
		}

		analytic.ParseAt = entity.Optional[time.Time]{Value: parseAt.Time, Set: parseAt.Valid}
		analytic.VacanciesNum = entity.Optional[int]{Value: int(vacanciesNum.Int64), Set: vacanciesNum.Valid}
		
		analyticResp = append(analyticResp, analytic)
	}

	if len(analyticResp) == 0 {
		return nil, repoerr.ErrAnalyticNotFound
	}

	return analyticResp, err
}
