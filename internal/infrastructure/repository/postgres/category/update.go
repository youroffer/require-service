package CategoryRepo

// import (
// 	"context"
// 	"errors"

// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/repoerr"
// 	"github.com/jackc/pgx/v5"
// )

// func (r *CategoryRepo) Update(ctx context.Context, category, title string) (*entity.Category, error) {
// 	newCategory := &entity.Category{}

// 	err := r.DB.QueryRow(ctx, `
// 	update categories
// 		set title = $2
// 	where title = $1
// 	returning *`, category, title).Scan(&newCategory.ID, &newCategory.Title)

// 	if errors.Is(err, pgx.ErrNoRows) {
// 		return nil, repoerr.ErrCategoryNotFound
// 	}

// 	return newCategory, err
// }
