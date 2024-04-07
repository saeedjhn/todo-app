package memorystore

import "github.com/saeedjhn/todo-app/domain"

type UserMemoryRepository struct {
}

func (u UserMemoryRepository) Save(user domain.User) error {
	//TODO implement me
	panic("implement me")
}

func (u UserMemoryRepository) Load() ([]domain.User, error) {
	//TODO implement me
	panic("implement me")
}
