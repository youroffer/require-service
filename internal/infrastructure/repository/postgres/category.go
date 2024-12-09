package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CategoryRepo struct {
	DB *pgxpool.Pool
}

func NewCategoryRepo(db *pgxpool.Pool) *CategoryRepo {
	return &CategoryRepo{DB: db}
}

func (r *CategoryRepo) GetAllWithPosts(ctx context.Context, public bool) (map[string][]*entity.PostResponse, error) {
	query := `
	SELECT 
		c.title AS category_title,
		p.id AS post_id, p.title AS post_title,
		p.public, p.logo_id
	FROM categories c
	INNER JOIN posts p ON c.id = p.categories_id
	%v
	ORDER BY c.title, p.title`

	filter := ""
	if public {
		filter = "WHERE p.public = true"
	}
	query = fmt.Sprintf(query, filter)

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make(map[string][]*entity.PostResponse)
	for rows.Next() {
		c := &entity.Category{}
		p := &entity.PostResponse{}

		if err := rows.Scan(
			&c.Title,
			&p.ID, &p.Title,
			&p.Public, &p.LogoID); err != nil {
			return nil, err
		}

		// Добавление категории
		if _, exists := categories[c.Title]; !exists {
			categories[c.Title] = []*entity.PostResponse{}
		}

		// Добавление должностей
		categories[c.Title] = append(categories[c.Title], p)
	}

	if len(categories) == 0 {
		return nil, repoerr.ErrPostNotFound
	}

	return categories, nil
}

func (r *CategoryRepo) GetAll(ctx context.Context) ([]*entity.Category, error) {
	var categories []*entity.Category
	rows, err := r.DB.Query(ctx, `select * from categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &entity.Category{}
		if err := rows.Scan(&c.ID, &c.Title); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if len(categories) == 0 {
		return nil, repoerr.ErrCategoryNotFound
	}

	return categories, nil
}

func (r *CategoryRepo) Update(ctx context.Context, category, title string) (*entity.Category, error) {
	newCategory := &entity.Category{}

	err := r.DB.QueryRow(ctx, `
	update categories 
		set title = $2 
	where title = $1
	returning *`, category, title).Scan(&newCategory.ID, &newCategory.Title)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repoerr.ErrCategoryNotFound
	}

	return newCategory, err
}

func (r *CategoryRepo) Delete(ctx context.Context, category string) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from categories where title = $1;`, category)
	if cmdTag.RowsAffected() == 0 {
		return repoerr.ErrCategoryNotFound
	}
	return err
}
