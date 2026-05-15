package user

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

func (s *Service) Get(ctx context.Context, uuid string) (entity.User, error) {
	user, err := s.userRepository.Get(ctx, uuid)
	if err != nil {
		return entity.User{}, entity.ErrGetUser
	}

	return user, nil
}
