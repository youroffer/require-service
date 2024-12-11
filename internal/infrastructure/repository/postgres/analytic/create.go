package analyticRepo

// import (
// 	"context"
// 	"errors"

// 	"github.com/Masterminds/squirrel"
// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
// 	"github.com/jackc/pgx/v5/pgconn"
// )

// func (r *AnalyticRepo) Create(ctx context.Context, qe repository.Querier, analytic *entity.Analytic) (*entity.AnalyticResp, error) {
// 	query, args, err := squirrel.
// 		Insert("analytics").
// 		Columns("post_id", "search_query").
// 		Values(analytic.PostID, analytic.SearchQuery).
// 		Suffix("returning id, word").
// 		PlaceholderFormat(squirrel.Dollar).
// 		ToSql()
// 	if err != nil {
// 		return nil, err
// 	}

// 	var id int
// 	err = qe.QueryRow(ctx, query, args...).Scan(&id)

// 	var pgErr *pgconn.PgError
// 	if err != nil {
// 		if errors.As(err, &pgErr) {
// 			if pgErr.Code == repoerr.FKViolation {
// 				return nil, repoerr.ErrAnalyticDependencyNotFound
// 			}
// 		}
// 		return nil, err
// 	}

// 	return r.GetByID(ctx, qe, id)
// }
