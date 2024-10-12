package postgres

import (
	"context"
	"errors"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) AddFilter(ctx context.Context, filter string) (*models.Filter, error) {
	newFilter := &models.Filter{}

	err := r.DB.QueryRow(ctx, `
	insert into filters 
		(word) 
	values 
		($1) 
	returning *;`, filter).Scan(
		&newFilter.ID, &newFilter.Word)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repository.UniqueConstraint {
			return nil, repository.ErrFilterExist
		}
	}

	return newFilter, err
}

func (r *Repository) DeleteFilter(ctx context.Context, filter string) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from filters where word = $1`, filter)
	if cmdTag.RowsAffected() == 0 {
		return repository.ErrFilterNotFound
	}
	return err
}

func (r *Repository) GetFilters(ctx context.Context) ([]*models.Filter, error) {
	rows, err := r.DB.Query(ctx, `select * from filters`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filters := []*models.Filter{}
	for rows.Next() {
		filter := &models.Filter{}
		if err := rows.Scan(&filter.ID, &filter.Word); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

func (r *Repository) GetFiltersWithPagination(ctx context.Context, limit, offset int) ([]*models.Filter, error) {
	rows, err := r.DB.Query(ctx, `
	select * 
	from filters
	limit $1 offset $2`, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	filters := []*models.Filter{}
	for rows.Next() {
		filter := &models.Filter{}
		if err := rows.Scan(&filter.ID, &filter.Word); err != nil {
			return nil, err
		}
		filters = append(filters, filter)
	}

	return filters, nil
}

func (r *Repository) GetFiltersCount(ctx context.Context) (int, error) {
	var count int
	err := r.DB.QueryRow(ctx, `SELECT COUNT(*) FROM filters;`).Scan(&count)
	return count, err
}
