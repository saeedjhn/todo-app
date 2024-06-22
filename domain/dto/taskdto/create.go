package taskdto

import "github.com/saeedjhn/todo-app/domain/entity"

type CreateRequest struct {
	AuthenticatedUserId int
	CategoryId          int
	Title               string
	DueDate             string
}

type CreateResponse struct {
	Task entity.Task
}
