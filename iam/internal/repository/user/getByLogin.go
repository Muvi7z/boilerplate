package user

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

func (r *Repository) GetByLogin(ctx context.Context, login string) (entity.User, error) {

	return entity.User{}, nil
}
