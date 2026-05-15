package iam

import (
	"context"
	"time"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

type sessionRepository interface {
	Get(ctx context.Context, uuid string) (entity.Session, error)
	Delete(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value entity.Session, ttl time.Duration) error
}

type userService interface {
	Get(ctx context.Context, uuid string) (entity.User, error)
	GetByLogin(ctx context.Context, login string) (entity.User, error)
	Create(ctx context.Context, user entity.User) (string, error)
}
