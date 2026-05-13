package user

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

type userRepository interface {
	Get(ctx context.Context, uuid string) (entity.User, error)
	Create(ctx context.Context, user *entity.User) (string, error)
}
