package iam

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Register(ctx context.Context, user entity.User) (string, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", entity.ErrRegisterUser
	}

	user.Password = string(passwordHash)
	userUUID := uuid.New().String()
	user.Uuid = userUUID

	createUUID, err := s.userService.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return createUUID, nil
}
