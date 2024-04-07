package domain

type Task struct {
	Id         int
	UserId     int
	CategoryId int
	Title      string
	DueDate    string
	IsDone     bool
}
