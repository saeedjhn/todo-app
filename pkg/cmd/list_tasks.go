package cmd

import (
	"fmt"
	"github.com/saeedjhn/todo-app/domain/dto/taskdto"
	"github.com/saeedjhn/todo-app/service/taskservice"
)

func ListTasks(ts taskservice.TaskAdaptor) {
	tasks, err := ts.List(taskdto.ListRequest{AuthenticatedUserId: AuthenticatedUser.Id})
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println("User tasks:", tasks.Tasks)
}
