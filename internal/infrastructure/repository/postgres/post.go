package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostRepo struct {
	DB *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *PostRepo {
	return &PostRepo{DB: db}
}

func (r *PostRepo) Add(ctx context.Context, post *entity.Position) (*entity.PositionResp, error) {
	postResp := &entity.PositionResp{}

	err := r.DB.QueryRow(ctx, `
	insert into posts 
		(categories_id, logo_id, title, public) 
	values 
		($1, $2, $3, $4) 
	returning id, logo_id, title, public;`, post.CategoriesID, post.LogoID, post.Title, post.Public).Scan(
		&postResp.ID, &postResp.LogoID, &postResp.Title, &postResp.Public)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repoerr.FKViolation:
			return nil, repoerr.ErrPostDependencyNotFound
		case repoerr.UniqueConstraint:
			return nil, repoerr.ErrPostExists
		}
	}

	return postResp, err
}

func (r *PostRepo) Update(ctx context.Context, id int, post *entity.PositionUpdate) (*entity.PositionResp, error) {
	var keys []string
	var values []interface{}

	if post.CategoriesID.Set {
		keys = append(keys, "categories_id=$1")
		values = append(values, post.CategoriesID.Value)
	}

	if post.LogoID.Set{
		keys = append(keys, fmt.Sprintf("logo_id=$%d", len(keys)+1))
		values = append(values, post.LogoID.Value)
	}

	if post.Title.Set {
		keys = append(keys, fmt.Sprintf("title=$%d", len(keys)+1))
		values = append(values, post.Title.Value)
	}

	if post.Public.Set {
		keys = append(keys, fmt.Sprintf("public=$%d", len(keys)+1))
		values = append(values, post.Public.Value)
	}

	values = append(values, id)
	query := fmt.Sprintf(`
	update posts set %s 
		where id = $%d
	returning id, logo_id, title, public`, strings.Join(keys, ", "), len(values))

	postResp := &entity.PositionResp{}
	err := r.DB.QueryRow(ctx, query, values...).Scan(
		&postResp.ID, &postResp.LogoID, &postResp.Title, &postResp.Public)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrPostNotFound
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case repoerr.FKViolation:
			return nil, repoerr.ErrPostDependencyNotFound
		case repoerr.UniqueConstraint:
			return nil, repoerr.ErrPostExists
		}
	}

	return postResp, err
}

func (r *PostRepo) Delete(ctx context.Context, id int) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from posts where id = $1`, id)

	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrPostNotFound
	}

	return err
}
