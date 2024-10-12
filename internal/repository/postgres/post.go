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

func (r *Repository) AddPost(ctx context.Context, post *models.Post) (*models.PostResponse, error) {
	postResp := &models.PostResponse{}

	err := r.DB.QueryRow(ctx, `
	insert into posts 
		(categories_id, logo_id, title, public) 
	values 
		($1, $2, $3, $4) 
	returning id, logo_id, title, public;`, post.CategoriesID, post.LogosID, post.Title, post.Public).Scan(
		&postResp.ID, &postResp.LogoID, &postResp.Title, &postResp.Public)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repository.FKViolation:
			return nil, repository.ErrPostDependencyNotFound
		case repository.UniqueConstraint:
			return nil, repository.ErrPostExists
		}
	}

	return postResp, err
}

func (r *Repository) UpdatePost(ctx context.Context, id int, post *models.PostUpdate) (*models.PostResponse, error) {
	var keys []string
	var values []interface{}

	if post.CategoriesID != nil {
		keys = append(keys, "categories_id=$1")
		values = append(values, post.CategoriesID)
	}

	if post.LogosID != nil {
		keys = append(keys, fmt.Sprintf("logo_id=$%d", len(keys)+1))
		values = append(values, post.LogosID)
	}

	if post.Title != nil {
		keys = append(keys, fmt.Sprintf("title=$%d", len(keys)+1))
		values = append(values, post.Title)
	}

	if post.Public != nil {
		keys = append(keys, fmt.Sprintf("public=$%d", len(keys)+1))
		values = append(values, post.Public)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update posts set %s 
		where id = $%d
	returning id, logo_id, title, public`, strings.Join(keys, ", "), len(values))

	postResp := &models.PostResponse{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(
		&postResp.ID, &postResp.LogoID, &postResp.Title, &postResp.Public)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrPostNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repository.FKViolation:
			return nil, repository.ErrPostDependencyNotFound
		case repository.UniqueConstraint:
			return nil, repository.ErrPostExists
		}
	}

	return postResp, err
}

func (r *Repository) DeletePost(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from posts where id = $1`, id)

	if cmdTag.RowsAffected() == 0 {
		return repository.ErrPostNotFound
	}

	return err
}
