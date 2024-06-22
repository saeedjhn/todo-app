package cmd

import (
	"bufio"
	"fmt"
	"github.com/saeedjhn/todo-app/domain/dto/taskdto"
	"github.com/saeedjhn/todo-app/service/taskservice"
	"os"
	"strconv"
)

func CreateTask(ts taskservice.TaskAdaptor) {
	fmt.Println("***** Create Task ******")
	var title, category, duedate string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the task title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category:")
	scanner.Scan()
	category = scanner.Text()

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category id is not valid integer, %v", err)
		return
	}

	fmt.Println("please enter the task due date:")
	scanner.Scan()
	duedate = scanner.Text()

	task, err := ts.Create(taskdto.CreateRequest{
		AuthenticatedUserId: AuthenticatedUser.Id,
		CategoryId:          categoryId,
		Title:               title,
		DueDate:             duedate,
	})
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println("Create task:", task)
}
