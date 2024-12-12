package postgres

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AnalyticRepo struct {
	DB *pgxpool.Pool
}

func NewAnalyticRepo(db *pgxpool.Pool) *AnalyticRepo {
	return &AnalyticRepo{DB: db}
}

func (r *AnalyticRepo) Add(ctx context.Context, analytic *entity.Analytic) (*entity.Analytic, error) {
	newAnalytic := &entity.Analytic{}
	err := r.DB.QueryRow(ctx, `
	insert into analytics 
		(posts_id, search_query) 
	values ($1, $2) 
	returning *;`, analytic.PostID, analytic.SearchQuery).Scan(
		&newAnalytic.ID, &newAnalytic.PostID, &newAnalytic.SearchQuery)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repoerr.FKViolation:
			return nil, repoerr.ErrAnalyticDependencyNotFound
		case repoerr.UniqueConstraint:
			return nil, repoerr.ErrAnalyticExist
		}
	}

	return newAnalytic, err
}

// func (r *AnalyticRepo) Update(ctx context.Context, id int, analytics *entity.AnalyticUpdate) (*entity.Analytic, error) {
// 	var keys []string
// 	var values []interface{}

// 	if analytics.PostID != nil {
// 		keys = append(keys, "posts_id=$1")
// 		values = append(values, analytics.PostID)
// 	}

// 	if analytics.SearchQuery != nil {
// 		keys = append(keys, fmt.Sprintf("search_query=$%d", len(values)+1))
// 		values = append(values, analytics.SearchQuery)
// 	}

// 	values = append(values, id)
// 	query := fmt.Sprintf(`
// 	update analytics
// 	set %v
// 	where id=$%d
// 	returning *;`, strings.Join(keys, ", "), len(values))

// 	newAnalytic := &entity.Analytic{}
// 	err := r.DB.QueryRow(ctx, query, values...).Scan(&newAnalytic.ID, &newAnalytic.PostID, &newAnalytic.SearchQuery)
// 	if errors.Is(err, pgx.ErrNoRows) {
// 		return nil, repoerr.ErrAnalyticNotFound
// 	}

// 	var pgErr *pgconn.PgError
// 	if errors.As(err, &pgErr) {
// 		switch pgErr.Code {
// 		case repoerr.FKViolation:
// 			return nil, repoerr.ErrAnalyticDependencyNotFound
// 		case repoerr.UniqueConstraint:
// 			return nil, repoerr.ErrPostIDExist
// 		}
// 	}

// 	return newAnalytic, err
// }

func (r *AnalyticRepo) Delete(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from analytics where id = $1`, id)
	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrAnalyticNotFound
	}

	return err
}

func (r *AnalyticRepo) GetPostID(ctx context.Context, analyticId int) (int, error) {
	var postId int
	err := r.DB.QueryRow(ctx, `select a.posts_id from analytics a where a.id = $1`, analyticId).Scan(&postId)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, repoerr.ErrAnalyticNotFound
	}

	return postId, err
}

func (r *AnalyticRepo) Get(ctx context.Context, postID int) (*entity.AnalyticResp, error) {
	analytic := &entity.AnalyticResp{}
	err := r.DB.QueryRow(ctx,
		`select id, search_query from analytics where posts_id = $1`,
		postID).Scan(&analytic.ID, &analytic.SearchQuery)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrAnalyticNotFound
	}
	return analytic, err
}

func (r *AnalyticRepo) GetAll(ctx context.Context) ([]*entity.Analytic, error) {
	rows, err := r.DB.Query(ctx, `select * from analytics`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	analytics := []*entity.Analytic{}
	for rows.Next() {
		analytic := &entity.Analytic{}
		if err := rows.Scan(
			&analytic.ID, &analytic.PostID, &analytic.SearchQuery); err != nil {
			return nil, err
		}

		analytics = append(analytics, analytic)
	}

	return analytics, err
}
