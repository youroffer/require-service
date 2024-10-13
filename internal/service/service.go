package service

import (
	"context"

	"github.com/himmel520/uoffer/require/models"
	"github.com/sirupsen/logrus"
)

//go:generate mockery --name=Cache --with-expecter
//go:generate mockery --name=Repository --with-expecter

type (
	Repository interface {
		Analytic
		Post
		Category
		Filter
	}

	Post interface {
		AddPost(ctx context.Context, post *models.Post) (*models.PostResponse, error)
		UpdatePost(ctx context.Context, id int, post *models.PostUpdate) (*models.PostResponse, error)
		DeletePost(ctx context.Context, id int) error
	}

	Category interface {
		GetCategoriesWithPosts(ctx context.Context, public bool) (map[string][]*models.PostResponse, error)
		GetAllCategories(ctx context.Context) ([]*models.Category, error)
		AddCategory(ctx context.Context, category *models.Category) (*models.Category, error)
		UpdateCategory(ctx context.Context, category, title string) (*models.Category, error)
		DeleteCategory(ctx context.Context, category string) error
	}

	Analytic interface {
		AddAnalytic(ctx context.Context, analytic *models.Analytic) (*models.Analytic, error)
		UpdateAnalytic(ctx context.Context, id int, analytics *models.AnalyticUpdate) (*models.Analytic, error)
		DeleteAnalytic(ctx context.Context, id int) error
		GetAnalytic(ctx context.Context, postID int) (*models.AnalyticResp, error)
		GetPostIDByAnalytic(ctx context.Context, id int) (int, error)
	}

	Filter interface {
		AddFilter(ctx context.Context, filter string) (*models.Filter, error)
		DeleteFilter(ctx context.Context, filter string) error
		GetFilters(ctx context.Context) ([]*models.Filter, error)
		GetFiltersWithPagination(ctx context.Context, limit, offset int) ([]*models.Filter, error)
		GetFiltersCount(ctx context.Context) (int, error)
	}

	Cache interface {
		DeleteCategoriesWithPublicPosts(ctx context.Context) error
		GetCategoriesWithPublicPosts(ctx context.Context) (map[string][]*models.PostResponse, error)
		SetCategoriesWithPublicPosts(ctx context.Context, categories map[string][]*models.PostResponse) error

		DeleteAnalyticWithWords(ctx context.Context, postID int) error
		SetAnalyticWithWords(ctx context.Context, data *models.AnalyticWithWords, postID int) error
		GetAnalyticWithWords(ctx context.Context, postID int) (*models.AnalyticWithWords, error)
	}

	Service struct {
		mediaUrl string
		repo     Repository
		cache    Cache
		log      *logrus.Logger
	}
)

func New(repo Repository, cache Cache, mediaUrl string, log *logrus.Logger) *Service {
	return &Service{
		mediaUrl: mediaUrl,
		repo:     repo,
		cache:    cache,
		log:      log,
	}
}
