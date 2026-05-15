package iam

import (
	"context"
	"github.com/Muvi7z/boilerplate/iam/internal/entity"
)

func (s *service) Whoami(ctx context.Context, sessionUUID string) (entity.User, error) {
	session, err := s.sessionRepository.Get(ctx, sessionUUID)
	if err != nil {
		return entity.User{}, err
	}

	user, err := s.userService.Get(ctx, session.UserId)
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
