package iam

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (s *service) Login(ctx context.Context, req entity.User) (string, error) {
	user, err := s.userService.GetByLogin(ctx, req.Login)
	if err != nil {
		//log TODO
		return "", entity.ErrInvalidCredentials
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		// log TODO
		return "", entity.ErrInvalidCredentials
	}

	sessionUuid := uuid.New().String()

	session := entity.Session{
		Uuid:   sessionUuid,
		UserId: user.Uuid,
	}

	err = s.sessionRepository.Set(ctx, sessionUuid, session, s.cacheTTL)
	if err != nil {
		return "", err
	}

	return sessionUuid, nil
}
