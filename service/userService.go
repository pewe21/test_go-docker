package service

import (
	"github.com/pewe21/go-docker/repository"
	"github.com/pewe21/go-docker/schema"
)

type UserService interface {
	GetUserByID(id int) (*schema.User, error)
	CreateUser(user *schema.User) error
	UpdateUser(id int, user *schema.User) error
	DeleteUser(id int) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) GetUserByID(id int) (*schema.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *userService) CreateUser(user *schema.User) error {
	return s.userRepository.Create(user)
}

func (s *userService) UpdateUser(id int, user *schema.User) error {
	u, err := s.userRepository.GetByID(id)
	if err != nil {
		return err
	}
	if u.ID != 0 {
		return s.userRepository.Update(u.ID, user)
	}
	return nil
}

func (s *userService) DeleteUser(id int) error {
	return s.userRepository.Delete(id)
}
