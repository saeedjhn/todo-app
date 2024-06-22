package memorystore

import (
	"fmt"
	"github.com/saeedjhn/todo-app/domain/entity"
)

type TaskRepository struct {
	tasks []entity.Task
}

func New() *TaskRepository {
	return &TaskRepository{tasks: make([]entity.Task, 0)}
}

func (tr *TaskRepository) DoesThisUserHaveThisCategoryId(authenticatedUserId, categoryId int) (bool, error) {
	for _, task := range tr.tasks {
		if task.UserId == authenticatedUserId && task.CategoryId == categoryId {
			return true, nil
		}
	}

	return false, fmt.Errorf("user dosen`t category in system")
}

func (tr *TaskRepository) Create(t entity.Task) (entity.Task, error) {
	t.Id = len(tr.tasks) + 1
	tr.tasks = append(tr.tasks, t)

	return t, nil
}

func (tr *TaskRepository) ListForUser(id int) ([]entity.Task, error) {
	var userTasks []entity.Task

	for _, task := range tr.tasks {
		if task.Id == id {
			userTasks = append(userTasks, task)
		}
	}

	return userTasks, nil
}
