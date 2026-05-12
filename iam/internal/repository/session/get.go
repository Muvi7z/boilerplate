package session

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/Muvi7z/boilerplate/iam/internal/repository/converter"
	entity2 "github.com/Muvi7z/boilerplate/iam/internal/repository/entity"
	"github.com/gomodule/redigo/redis"
)

func (r *Repository) Get(ctx context.Context, uuid string) (entity.Session, error) {
	cacheKey := r.getCacheKey(uuid)

	values, err := r.cache.HGetAll(ctx, cacheKey)
	if err != nil {
		return entity.Session{}, err
	}

	if len(values) == 0 {
		return entity.Session{}, entity.ErrSessionNotFound
	}

	var sessionView entity2.SessionRedisView

	err = redis.ScanStruct(values, &sessionView)
	if err != nil {
		return entity.Session{}, err
	}

	return converter.SessionFromRedisView(sessionView), nil
}
