package redis

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/himmel520/uoffer/require/internal/repository"
	"github.com/himmel520/uoffer/require/models"
	"github.com/redis/go-redis/v9"
)

func (r *Redis) GetAnalyticWithWords(ctx context.Context, postID int) (*models.AnalyticWithWords, error) {
	val, err := r.rdb.Get(ctx, fmt.Sprintf("require:analytic:%d", postID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, repository.ErrAnalyticNotFound
		}
		return nil, err
	}

	analytic := new(models.AnalyticWithWords)
	err = json.Unmarshal([]byte(val), analytic)

	return analytic, err
}

func (r *Redis) SetAnalyticWithWords(ctx context.Context, analytic *models.AnalyticWithWords, postID int) error {
	analyticByte, err := json.Marshal(analytic)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, fmt.Sprintf("require:analytic:%d", postID), string(analyticByte), r.expiration).Result()
	return err
}

func (r *Redis) DeleteAnalyticWithWords(ctx context.Context, postID int) error {
	_, err := r.rdb.Del(ctx, fmt.Sprintf("require:analytic:%d", postID)).Result()
	if err == redis.Nil {
		return nil
	}
	return err
}
