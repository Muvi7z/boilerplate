package iam

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

func (s *service) Get(ctx context.Context, uuid string) (entity.User, error) {
	user, err := s.userService.Get(ctx, uuid)
	if err != nil {
		return entity.User{}, entity.ErrGetUser
	}

	return user, nil
}
