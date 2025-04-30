package user

type Service struct {
	repo Repository
}

func NewUserService(repo Repository) *Service {
	return &Service{repo}
}

func (s *Service) CreateUser(u *User) error {
	return s.repo.Create(u)
}
