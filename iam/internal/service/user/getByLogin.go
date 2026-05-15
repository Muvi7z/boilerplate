package user

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

func (s *Service) GetByLogin(ctx context.Context, login string) (entity.User, error) {
	user, err := s.userRepository.GetByLogin(ctx, login)
	if err != nil {
		return entity.User{}, entity.ErrGetUser
	}

	return user, nil
}
