package session

import (
	"fmt"

	"github.com/Muvi7z/boilerplate/platform/cache"
)

const (
	cacheKeyPrefix = "iam:iam:"
)

type Repository struct {
	cache cache.RedisClient
}

func NewRepository(cache cache.RedisClient) *Repository {
	return &Repository{
		cache: cache,
	}
}

func (r *Repository) getCacheKey(uuid string) string {
	return fmt.Sprintf("%s%s", cacheKeyPrefix, uuid)
}
