package user

type Service struct {
	userRepository userRepository
}

func New(userRepository userRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
