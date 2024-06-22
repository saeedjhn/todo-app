package memorystore

import (
	"github.com/saeedjhn/todo-app/domain/entity"
)

type UserRepository struct {
}

func (u *UserRepository) Save(user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserRepository) Load() ([]entity.User, error) {
	//TODO implement me
	panic("implement me")
}
