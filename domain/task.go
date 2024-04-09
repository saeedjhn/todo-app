package domain

type Task struct {
	Id         int
	UserId     int
	CategoryId int
	Title      string
	DueDate    string
	IsDone     bool
}

type TaskAdaptorI interface {
	Create(tr TaskCreateRequest) (TaskCreateResponse, error)
	List(tr TaskListRequest) (TaskListResponse, error)
}

type TaskRepositoryI interface {
	DoesThisUserHaveThisCategoryId(authenticatedUserId, categoryId int) (bool, error)
	CreateNewTask(t Task) (Task, error)
	ListUserTasks(id int) ([]Task, error)
}

/* Data Transfer Object - DTO */

type TaskCreateRequest struct {
	AuthenticatedUserId int
	CategoryId          int
	Title               string
	DueDate             string
}

type TaskCreateResponse struct {
	Task Task
}

type TaskListRequest struct {
	AuthenticatedUserId int
}

type TaskListResponse struct {
	Tasks []Task
}
