package user

type service struct {
	userRepository userRepository
}

func New(userRepository userRepository) *service {
	return &service{
		userRepository: userRepository,
	}
}
