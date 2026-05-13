package iam

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/google/uuid"
)

func (s *service) Login(ctx context.Context, user entity.User) (string, error) {

	s.userService.Get(ctx, user.Uuid)

	sessionUuid := uuid.New().String()

	session := entity.Session{
		Uuid:   sessionUuid,
		UserId: user.Uuid,
	}

	err := s.sessionRepository.Set(ctx, sessionUuid, session, s.cacheTTL)
	if err != nil {
		return "", err
	}

	return sessionUuid, nil
}
