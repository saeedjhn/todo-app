package service

import "github.com/saeedjhn/todo-app/domain"

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(userRepository domain.UserRepository) domain.UserAdaptorI {
	return UserService{userRepository: userRepository}
}

func (u UserService) Save(user domain.User) error {
	return u.userRepository.Save(user)
}

func (u UserService) Load() ([]domain.User, error) {
	return u.userRepository.Load()
}
