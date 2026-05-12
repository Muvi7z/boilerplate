package session

import (
	"context"
	"time"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/Muvi7z/boilerplate/iam/internal/repository/converter"
)

func (r *Repository) Set(ctx context.Context, key string, value entity.Session, ttl time.Duration) error {
	cacheKey := r.getCacheKey(key)

	redisView := converter.SessionToRedisView(value)

	err := r.cache.HashSet(ctx, cacheKey, redisView)
	if err != nil {
		return err
	}
	return r.cache.Expire(ctx, cacheKey, ttl)
}
