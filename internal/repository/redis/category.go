package redis

import (
	"context"
	"encoding/json"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) GetCategoriesWithPublicPosts(ctx context.Context) (map[string][]*models.PostResponse, error) {
	val, err := r.rdb.Get(ctx, "require:categories:posts").Result()
	if err != nil {
		if err == redis.Nil {
			return nil, repository.ErrKeyNotFound
		}
		return nil, err
	}

	categories := new(map[string][]*models.PostResponse)
	err = json.Unmarshal([]byte(val), categories)

	return *categories, err
}

func (r *Redis) SetCategoriesWithPublicPosts(ctx context.Context, categories map[string][]*models.PostResponse) error {
	categoriesByte, err := json.Marshal(categories)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, "require:categories:posts", string(categoriesByte), r.expiration).Result()
	return err
}

func (r *Redis) DeleteCategoriesWithPublicPosts(ctx context.Context) error {
	_, err := r.rdb.Del(ctx, "require:categories:posts").Result()
	if err == redis.Nil {
		return nil
	}
	return err
}
