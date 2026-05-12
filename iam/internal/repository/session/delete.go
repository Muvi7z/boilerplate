package session

import "context"

func (r *Repository) Delete(ctx context.Context, key string) error {
	cacheKey := r.getCacheKey(key)
	return r.cache.Del(ctx, cacheKey)
}
