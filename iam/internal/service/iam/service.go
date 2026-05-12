package iam

type service struct {
	sessionRepository sessionRepository
}

func New(sessionRepository sessionRepository) *service {
	return &service{
		sessionRepository: sessionRepository,
	}
}
