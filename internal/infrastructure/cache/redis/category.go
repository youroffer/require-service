package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache/errcache"

	"github.com/redis/go-redis/v9"
)

type CategoryCache struct {
	rdb *redis.Client
	exp time.Duration
}

func NewCategoryCache(rdb *redis.Client, exp time.Duration) *CategoryCache {
	return &CategoryCache{rdb: rdb, exp: exp}
}

func (r *CategoryCache) GetAllWithPublicPosts(ctx context.Context) (map[string][]*entity.PostResponse, error) {
	val, err := r.rdb.Get(ctx, "require:categories:posts").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errcache.ErrKeyNotFound
		}
		return nil, err
	}

	categories := new(map[string][]*entity.PostResponse)
	err = json.Unmarshal([]byte(val), categories)

	return *categories, err
}

func (r *CategoryCache) SetAllWithPublicPosts(ctx context.Context, categories map[string][]*entity.PostResponse) error {
	categoriesByte, err := json.Marshal(categories)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, "require:categories:posts", string(categoriesByte), r.exp).Result()
	return err
}

func (r *CategoryCache) DeleteAllWithPublicPosts(ctx context.Context) error {
	_, err := r.rdb.Del(ctx, "require:categories:posts").Result()
	if err == redis.Nil {
		return nil
	}
	return err
}
