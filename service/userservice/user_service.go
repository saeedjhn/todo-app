package userservice

import (
	"github.com/saeedjhn/todo-app/domain/dto/userdto"
	"github.com/saeedjhn/todo-app/domain/entity"
)

type UserAdaptor interface {
	Save(s userdto.SaveRequest) error
	Load() (userdto.LoadResponse, error)
}

type Repository interface {
	Save(u entity.User) error
	Load() ([]entity.User, error)
}

type UserService struct {
	repo Repository
}

func New(repo Repository) *UserService {
	return &UserService{repo: repo}
}

func (u UserService) Save(s userdto.SaveRequest) error {
	user := entity.User{
		Email:    s.Email,
		Password: s.Password,
	}

	return u.repo.Save(user)
}

func (u UserService) Load() (userdto.LoadResponse, error) {
	users, err := u.repo.Load()

	return userdto.LoadResponse{Users: users}, err
}
