package repository

import (
	"context"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/repository/postgres"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Repository struct {
		Post     PostRepo
		Category CategoryRepo
		Analytic AnalyticRepo
		Filter   FilterRepo
	}

	PostRepo interface {
		Add(ctx context.Context, post *entity.Post) (*entity.PostResponse, error)
		Update(ctx context.Context, id int, post *entity.PostUpdate) (*entity.PostResponse, error)
		Delete(ctx context.Context, id int) error
	}

	CategoryRepo interface {
		GetAllWithPosts(ctx context.Context, public bool) (map[string][]*entity.PostResponse, error)
		GetAll(ctx context.Context) ([]*entity.Category, error)
		Add(ctx context.Context, category *entity.Category) (*entity.Category, error)
		Update(ctx context.Context, category, title string) (*entity.Category, error)
		Delete(ctx context.Context, category string) error
	}

	AnalyticRepo interface {
		Add(ctx context.Context, analytic *entity.Analytic) (*entity.Analytic, error)
		Update(ctx context.Context, id int, analytics *entity.AnalyticUpdate) (*entity.Analytic, error)
		Delete(ctx context.Context, id int) error
		Get(ctx context.Context, postID int) (*entity.AnalyticResp, error)
		GetPostID(ctx context.Context, analyticId int) (int, error)
		GetAll(ctx context.Context) ([]*entity.Analytic, error)
	}

	FilterRepo interface {
		Add(ctx context.Context, filter string) (*entity.Filter, error)
		Delete(ctx context.Context, filter string) error
		GetAll(ctx context.Context) ([]*entity.Filter, error)
		GetWithPagination(ctx context.Context, limit, offset int) ([]*entity.Filter, error)
		GetCount(ctx context.Context) (int, error)
	}
)

func New(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Post:     postgres.NewPostRepo(pool),
		Category: postgres.NewCategoryRepo(pool),
		Analytic: postgres.NewAnalyticRepo(pool),
		Filter:   postgres.NewFilterRepo(pool),
	}
}
