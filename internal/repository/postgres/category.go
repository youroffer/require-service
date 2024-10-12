package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

func (r *Repository) GetCategoriesWithPosts(ctx context.Context, public bool) (map[string][]*models.PostResponse, error) {
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

	categories := make(map[string][]*models.PostResponse)
	for rows.Next() {
		c := &models.Category{}
		p := &models.PostResponse{}

		if err := rows.Scan(
			&c.Title,
			&p.ID, &p.Title,
			&p.Public, &p.LogoID); err != nil {
			return nil, err
		}

		// Добавление категории
		if _, exists := categories[c.Title]; !exists {
			categories[c.Title] = []*models.PostResponse{}
		}

		// Добавление должностей
		categories[c.Title] = append(categories[c.Title], p)
	}

	if len(categories) == 0 {
		return nil, repository.ErrPostNotFound
	}

	return categories, nil
}

func (r *Repository) GetAllCategories(ctx context.Context) ([]*models.Category, error) {
	var categories []*models.Category
	rows, err := r.DB.Query(ctx, `select * from categories`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c := &models.Category{}
		if err := rows.Scan(&c.ID, &c.Title); err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	if len(categories) == 0 {
		return nil, repository.ErrCategoryNotFound
	}

	return categories, nil
}

func (r *Repository) AddCategory(ctx context.Context, category *models.Category) (*models.Category, error) {
	newCategory := &models.Category{}

	err := r.DB.QueryRow(ctx, `
	insert into categories 
		(title) 
	values 
		($1) 
	returning *`, category.Title).Scan(&newCategory.ID, &newCategory.Title)

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		if pgErr.Code == repository.UniqueConstraint {
			return nil, repository.ErrCategoryExists
		}
	}

	return newCategory, err
}

func (r *Repository) UpdateCategory(ctx context.Context, category, title string) (*models.Category, error) {
	newCategory := &models.Category{}

	err := r.DB.QueryRow(ctx, `
	update categories 
		set title = $2 
	where title = $1
	returning *`, category, title).Scan(&newCategory.ID, &newCategory.Title)

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, repository.ErrCategoryNotFound
	}

	return newCategory, err
}

func (r *Repository) DeleteCategory(ctx context.Context, category string) error {
	cmdTag, err := r.DB.Exec(ctx, `delete from categories where title = $1;`, category)
	if cmdTag.RowsAffected() == 0 {
		return repository.ErrCategoryNotFound
	}
	return err
}
