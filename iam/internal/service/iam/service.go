package iam

import (
	"time"
)

type Service struct {
	sessionRepository sessionRepository
	userService       userService
	cacheTTL          time.Duration
}

func New(sessionRepository sessionRepository, userService userService) *Service {
	return &Service{
		sessionRepository: sessionRepository,
		userService:       userService,
	}
}
