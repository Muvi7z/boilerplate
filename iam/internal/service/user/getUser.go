package user

import (
	"context"

	"github.com/Muvi7z/boilerplate/iam/internal/entity"
	"github.com/Muvi7z/boilerplate/iam/internal/service/iam"
)

func (s *iam.service) GetUser(ctx context.Context, user entity.User) (string, error) {

	var session string

	s.sessionRepository.Get(ctx, user.Uuid)

	return session, nil
}
