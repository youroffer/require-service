package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/himmel520/uoffer/require/internal/entity"
	"github.com/himmel520/uoffer/require/internal/infrastructure/cache/errcache"

	"github.com/redis/go-redis/v9"
)

type AnalyticCache struct {
	rdb *redis.Client
	exp time.Duration
}

func NewAnalyticCache(rdb *redis.Client, exp time.Duration) *AnalyticCache {
	return &AnalyticCache{rdb: rdb, exp: exp}
}

func (r *AnalyticCache) GetWithWords(ctx context.Context, postID int) (*entity.AnalyticWithWords, error) {
	val, err := r.rdb.Get(ctx, fmt.Sprintf("require:analytic:%d", postID)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, errcache.ErrKeyNotFound
		}
		return nil, err
	}

	analytic := new(entity.AnalyticWithWords)
	err = json.Unmarshal([]byte(val), analytic)

	return analytic, err
}

func (r *AnalyticCache) SetWithWords(ctx context.Context, analytic *entity.AnalyticWithWords, postID int) error {
	analyticByte, err := json.Marshal(analytic)
	if err != nil {
		return err
	}

	_, err = r.rdb.Set(ctx, fmt.Sprintf("require:analytic:%d", postID), string(analyticByte), r.exp).Result()
	return err
}

func (r *AnalyticCache) DeleteWithWords(ctx context.Context, postID int) error {
	_, err := r.rdb.Del(ctx, fmt.Sprintf("require:analytic:%d", postID)).Result()
	if err == redis.Nil {
		return nil
	}
	return err
}
