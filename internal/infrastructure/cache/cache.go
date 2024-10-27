package cache

import (
	"context"
	"time"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache/redis"

	goredis "github.com/redis/go-redis/v9"
)

type (
	Cache struct {
		Analytic AnalyticCache
		Category CategoryCache
	}

	AnalyticCache interface {
		GetWithWords(ctx context.Context, postID int) (*entity.AnalyticWithWords, error)
		SetWithWords(ctx context.Context, analytic *entity.AnalyticWithWords, postID int) error
		DeleteWithWords(ctx context.Context, postID int) error
	}

	CategoryCache interface {
		GetAllWithPublicPosts(ctx context.Context) (map[string][]*entity.PostResponse, error)
		SetAllWithPublicPosts(ctx context.Context, categories map[string][]*entity.PostResponse) error
		DeleteAllWithPublicPosts(ctx context.Context) error
	}
)

func New(rdb *goredis.Client, exp time.Duration) *Cache {
	return &Cache{
		Analytic: redis.NewAnalyticCache(rdb, exp),
		Category: redis.NewCategoryCache(rdb, exp),
	}
}
