package memorystore

import (
	"fmt"
	"github.com/saeedjhn/todo-app/domain"
)

type TaskMemoryRepository struct {
	tasks []domain.Task
}

func NewTaskMemoryRepository() domain.TaskRepositoryI {
	return &TaskMemoryRepository{tasks: make([]domain.Task, 0)}
}

func (tr *TaskMemoryRepository) DoesThisUserHaveThisCategoryId(authenticatedUserId, categoryId int) (bool, error) {
	for _, task := range tr.tasks {
		if task.UserId == authenticatedUserId && task.CategoryId == categoryId {
			return true, nil
		}
	}

	return false, fmt.Errorf("user dosen`t category in system")
}

func (tr *TaskMemoryRepository) CreateNewTask(t domain.Task) (domain.Task, error) {
	t.Id = len(tr.tasks) + 1
	tr.tasks = append(tr.tasks, t)

	return t, nil
}

func (tr *TaskMemoryRepository) ListUserTasks(id int) ([]domain.Task, error) {
	var userTasks []domain.Task

	for _, task := range tr.tasks {
		if task.Id == id {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks, nil
}
