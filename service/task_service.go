package service

import (
	"fmt"
	"github.com/saeedjhn/todo-app/domain"
)

type TaskService struct {
	repository domain.TaskRepositoryI
}

func NewTaskService(repository domain.TaskRepositoryI) domain.TaskAdaptorI {
	return &TaskService{repository: repository}
}

func (ts TaskService) Create(tr domain.TaskCreateRequest) (domain.TaskCreateResponse, error) {
	// TODO: Create CategoryService & CategoryXRepository
	//ok, err := ts.repository.DoesThisUserHaveThisCategoryId(tr.AuthenticatedUserId, tr.CategoryId)
	//if err != nil {
	//	return domain.TaskCreateResponse{}, err
	//}
	//if !ok {
	//	return domain.TaskCreateResponse{}, fmt.Errorf("user does not have this category")
	//}

	task, err := ts.repository.CreateNewTask(domain.Task{
		UserId:     tr.AuthenticatedUserId,
		CategoryId: tr.CategoryId,
		Title:      tr.Title,
		DueDate:    tr.DueDate,
		IsDone:     false,
	})

	if err != nil {
		return domain.TaskCreateResponse{}, fmt.Errorf("can`t create task %s", err.Error())
	}

	return domain.TaskCreateResponse{Task: task}, nil
}

func (ts TaskService) List(tlr domain.TaskListRequest) (domain.TaskListResponse, error) {
	tasks, err := ts.repository.ListUserTasks(tlr.AuthenticatedUserId)
	if err != nil {
		return domain.TaskListResponse{}, fmt.Errorf("can't list user tasks: %v", err)
	}

	return domain.TaskListResponse{Tasks: tasks}, nil
}
