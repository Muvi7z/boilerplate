package iam

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/google/uuid"
)

func (s *service) Login(ctx context.Context, user entity.User) (string, error) {

	sessionUuid := uuid.New().String()

	session := entity.Session{
		Uuid:   sessionUuid,
		UserId: user.Uuid,
	}

	s.sessionRepository.Set(ctx, session)

	return session, nil
}
