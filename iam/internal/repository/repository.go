package repository

import "context"

type IAMCacheRepository interface {
	Get(ctx context.Context, uuid string)
}
