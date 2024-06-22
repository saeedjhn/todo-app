package taskservice

import (
	"fmt"
	"github.com/saeedjhn/todo-app/domain/dto/taskdto"
	"github.com/saeedjhn/todo-app/domain/entity"
)

type TaskAdaptor interface {
	Create(tr taskdto.CreateRequest) (taskdto.CreateResponse, error)
	List(tr taskdto.ListRequest) (taskdto.ListResponse, error)
}

type Repository interface {
	DoesThisUserHaveThisCategoryId(authenticatedUserId, categoryId int) (bool, error)
	Create(t entity.Task) (entity.Task, error)
	ListForUser(id int) ([]entity.Task, error)
}

type TaskService struct {
	repo Repository
}

func New(repository Repository) *TaskService {
	return &TaskService{repo: repository}
}

func (ts TaskService) Create(tr taskdto.CreateRequest) (taskdto.CreateResponse, error) {
	// TODO: Create CategoryService & CategoryXRepository
	//ok, err := ts.repo.DoesThisUserHaveThisCategoryId(tr.AuthenticatedUserId, tr.CategoryId)
	//if err != nil {
	//	return domain.TaskCreateResponse{}, err
	//}
	//if !ok {
	//	return domain.TaskCreateResponse{}, fmt.Errorf("user does not have this category")
	//}

	task, err := ts.repo.Create(entity.Task{
		UserId:     tr.AuthenticatedUserId,
		CategoryId: tr.CategoryId,
		Title:      tr.Title,
		DueDate:    tr.DueDate,
		IsDone:     false,
	})

	if err != nil {
		return taskdto.CreateResponse{}, fmt.Errorf("can`t create task %s", err.Error())
	}

	return taskdto.CreateResponse{Task: task}, nil
}

func (ts TaskService) List(tlr taskdto.ListRequest) (taskdto.ListResponse, error) {
	tasks, err := ts.repo.ListForUser(tlr.AuthenticatedUserId)
	if err != nil {
		return taskdto.ListResponse{}, fmt.Errorf("can't list user tasks: %v", err)
	}

	return taskdto.ListResponse{Tasks: tasks}, nil
}
