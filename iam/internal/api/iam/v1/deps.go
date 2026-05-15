package v1

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

type sessionService interface {
	Get(ctx context.Context, uuid string) (entity.User, error)
	Login(ctx context.Context, req entity.User) (string, error)
	Register(ctx context.Context, user entity.User) (string, error)
}
