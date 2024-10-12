package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) AddAnalytic(ctx context.Context, analytic *models.Analytic) (*models.Analytic, error) {
	newAnalytic := &models.Analytic{}
	err := r.DB.QueryRow(ctx, `
	insert into analytics 
		(posts_id, search_query) 
	values ($1, $2) 
	returning *;`, analytic.PostID, analytic.SearchQuery).Scan(
		&newAnalytic.ID, &newAnalytic.PostID, &newAnalytic.SearchQuery)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repository.FKViolation:
			return nil, repository.ErrAnalyticDependencyNotFound
		case repository.UniqueConstraint:
			return nil, repository.ErrPostIDExist
		}
	}

	return newAnalytic, err
}

func (r *Repository) UpdateAnalytic(ctx context.Context, id int, analytics *models.AnalyticUpdate) (*models.Analytic, error) {
	var keys []string
	var values []interface{}

	if analytics.PostID != nil {
		keys = append(keys, "posts_id=$1")
		values = append(values, analytics.PostID)
	}

	if analytics.SearchQuery != nil {
		keys = append(keys, fmt.Sprintf("search_query=$%d", len(values)+1))
		values = append(values, analytics.SearchQuery)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update analytics 
	set %v 
	where id=$%d
	returning *;`, strings.Join(keys, ", "), len(values))

	newAnalytic := &models.Analytic{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(&newAnalytic.ID, &newAnalytic.PostID, &newAnalytic.SearchQuery)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrAnalyticNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repository.FKViolation:
			return nil, repository.ErrAnalyticDependencyNotFound
		case repository.UniqueConstraint:
			return nil, repository.ErrPostIDExist
		}
	}

	return newAnalytic, err
}

func (r *Repository) DeleteAnalytic(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from analytics where id = $1`, id)
	if cmdTag.RowsAffected() == 0 {
		return repository.ErrAnalyticNotFound
	}

	return err
}

func (r *Repository) GetPostIDByAnalytic(ctx context.Context, id int) (int, error) {
	var postId int
	err := r.DB.QueryRow(ctx, `select a.posts_id from analytics a where a.id = $1`, id).Scan(&postId)
	if errors.Is(err, pgx.ErrNoRows) {
		return 0, repository.ErrAnalyticNotFound
	}

	return postId, err
}

func (r *Repository) GetAnalytic(ctx context.Context, postID int) (*models.AnalyticResp, error) {
	analytic := &models.AnalyticResp{}
	err := r.DB.QueryRow(ctx,
		`select id, search_query from analytics where posts_id = $1`,
		postID).Scan(&analytic.ID, &analytic.SearchQuery)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrAnalyticNotFound
	}
	return analytic, err
}

func (r *Repository) GetAnalytics(ctx context.Context) ([]*models.Analytic, error) {
	rows, err := r.DB.Query(ctx, `select * from analytics`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	analytics := []*models.Analytic{}
	for rows.Next() {
		analytic := &models.Analytic{}
		if err := rows.Scan(
			&analytic.ID, &analytic.PostID, &analytic.SearchQuery); err != nil {
			return nil, err
		}

		analytics = append(analytics, analytic)
	}

	return analytics, err
}
