package session

import (
	"context"
	"errors"
	"time"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/Muvi7z/boilerplate/iam/internal/repository/converter"
	"github.com/gomodule/redigo/redis"
)

func (r *repository) Set(ctx context.Context, key string, value entity.Session, ttl time.Duration) error {
	cacheKey := r.getCacheKey(key)

	redisView := converter.SessionToRedisView(value)

	//values, err := r.cache.HashSet(ctx, cacheKey, redisView)
	//if err != nil {
	//	return err
	//}
	return nil
}
