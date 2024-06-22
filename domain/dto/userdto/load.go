package userdto

import "github.com/saeedjhn/todo-app/domain/entity"

type LoadResponse struct {
	Users []entity.User
}
