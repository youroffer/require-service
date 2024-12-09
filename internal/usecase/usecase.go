package usecase

// import (
// 	"context"
// 	"crypto/rsa"

// 	"github.com/himmel520/uoffer/require/internal/entity"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/cache"
// 	"github.com/himmel520/uoffer/require/internal/infrastructure/repository"

// 	"github.com/sirupsen/logrus"
// )

// //go:generate mockery --all

// type (
// 	Usecase struct {
// 		Auth     AuthUC
// 		Analytic AnalyticUC
// 		Category CategoryUC
// 		Filter   FilterUC
// 		Post     PostUC
// 	}

// 	AnalyticUC interface {
// 		Add(ctx context.Context, analytic *entity.Analytic) (*entity.Analytic, error)
// 		Update(ctx context.Context, id int, analytic *entity.AnalyticUpdate) (*entity.Analytic, error)
// 		Delete(ctx context.Context, id int) error
// 		GetWithWords(ctx context.Context, postID int, role string) (*entity.AnalyticWithWords, error)
// 	}

// 	AuthUC interface {
// 		GetUserRoleFromToken(jwtToken string,) (string, error)
// 	}

// 	CategoryUC interface {
// 		GetAllWithPublicPosts(ctx context.Context) (map[string][]*entity.PostResponse, error)
// 		GetAllWithPosts(ctx context.Context) (map[string][]*entity.PostResponse, error)
// 		GetAll(ctx context.Context) ([]*entity.Category, error)
// 		Add(ctx context.Context, category *entity.Category) (*entity.Category, error)
// 		Update(ctx context.Context, category, title string) (*entity.Category, error)
// 		Delete(ctx context.Context, category string) error
// 		DeleteCache(ctx context.Context) error
// 	}

// 	FilterUC interface {
// 		Add(ctx context.Context, filter string) (*entity.Filter, error)
// 		Delete(ctx context.Context, filter string) error
// 		GetWithPagination(ctx context.Context, limit, offset int) (*entity.FilterResp, error)
// 	}

// 	PostUC interface {
// 		Add(ctx context.Context, post *entity.Post) (*entity.PostResponse, error)
// 		Update(ctx context.Context, id int, post *entity.PostUpdate) (*entity.PostResponse, error)
// 		Delete(ctx context.Context, id int) error
// 	}
// )

// func New(repo *repository.Repository, cache *cache.Cache, publicKey rsa.PublicKey, log *logrus.Logger) *Usecase {
// 	return &Usecase{
// 		Auth:     NewAuthUsecase(publicKey),
// 		Analytic: NewAnalyticUsecase(repo.Analytic, cache.Analytic, log),
// 		Category: NewCategoryUsecase(repo.Category, cache.Category, log),
// 		Filter:   NewFilterUsecase(repo.Filter, log),
// 		Post:     NewPostUsecase(repo.Post, log),
// 	}
// }
