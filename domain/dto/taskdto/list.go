package taskdto

import "github.com/saeedjhn/todo-app/domain/entity"

type ListRequest struct {
	AuthenticatedUserId int
}

type ListResponse struct {
	Tasks []entity.Task
}
