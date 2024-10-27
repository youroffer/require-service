package postgres

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type FilterRepo struct {
	DB *pgxpool.Pool
}

func NewFilterRepo(db *pgxpool.Pool) *FilterRepo {
	return &FilterRepo{DB: db}
}

func (r *FilterRepo) Add(ctx context.Context, filter string) (*entity.Filter, error) {
	newFilter := &entity.Filter{}

	err := r.DB.QueryRow(ctx, `
	insert into filters 
		(word) 
	values 
		($1) 
	returning *;`, filter).Scan(
		&newFilter.ID, &newFilter.Word)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repoerr.UniqueConstraint {
			return nil, repoerr.ErrFilterExist
		}
	}

	return newFilter, err
}

func (r *FilterRepo) Delete(ctx context.Context, filter string) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from filters where word = $1`, filter)
	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrFilterNotFound
	}
	return err
}

func (r *FilterRepo) GetAll(ctx context.Context) ([]*entity.Filter, error) {
	rows, err := r.DB.Query(ctx, `select * from filters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filters := []*entity.Filter{}
	for rows.Next() {
		filter := &entity.Filter{}
		if err := rows.Scan(&filter.ID, &filter.Word); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

func (r *FilterRepo) GetWithPagination(ctx context.Context, limit, offset int) ([]*entity.Filter, error) {
	rows, err := r.DB.Query(ctx, `
	select * 
	from filters
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filters := []*entity.Filter{}
	for rows.Next() {
		filter := &entity.Filter{}
		if err := rows.Scan(&filter.ID, &filter.Word); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

func (r *FilterRepo) GetCount(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM filters;`).Scan(&count)
	return count, err
}
