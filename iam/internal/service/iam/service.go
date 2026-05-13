package iam

import (
	"time"
)

type service struct {
	sessionRepository sessionRepository
	userService       userService
	cacheTTL          time.Duration
}

func New(sessionRepository sessionRepository) *service {
	return &service{
		sessionRepository: sessionRepository,
	}
}
