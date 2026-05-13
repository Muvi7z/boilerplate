package user

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/google/uuid"
)

func (s *service) Create(ctx context.Context, user entity.User) (string, error) {
	userUUID := uuid.New().String()

	user.Uuid = userUUID
	res, err := s.userRepository.Create(ctx, &user)
	if err != nil {
		return "", entity.ErrCreateUser
	}

	return res, nil
}
