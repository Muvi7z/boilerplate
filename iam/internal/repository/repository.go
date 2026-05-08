package repository

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"time"
)

type SessionRepository interface {
	Get(ctx context.Context, uuid string)
	Delete(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value entity.Session, ttl time.Duration) error
}
