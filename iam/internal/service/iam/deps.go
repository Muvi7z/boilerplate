package iam

import (
	"context"
	"time"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

type sessionRepository interface {
	Get(ctx context.Context, uuid string) entity.User
	Delete(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value entity.Session, ttl time.Duration) error
}
