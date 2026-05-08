package session

import "context"

func (r *repository) Delete(ctx context.Context, key string) error {
	cacheKey := r.getCacheKey(key)
	return r.cache.Del(ctx, cacheKey)
}
