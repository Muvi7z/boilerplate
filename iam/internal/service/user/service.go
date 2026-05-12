package user

type service struct {
}

func New(sessionRepository sessionRepository) *service {
	return &service{
		sessionRepository: sessionRepository,
	}
}
