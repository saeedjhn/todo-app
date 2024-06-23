package alarm

import (
	"fmt"
	"github.com/saeedjhn/todo-app/constant"
)

func AllCommands() string {

	return fmt.Sprintf("You are authorized to the following commands:\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n",
		constant.RegisterUser, constant.LoginUser, constant.CreateCategory, constant.CreateTask, constant.ListTask,
	)
}

func FieldCommand(command string) string {

	return fmt.Sprintf("Command is not valid: %s\n\nYou are authorized to the following commands:\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n",
		command, constant.RegisterUser, constant.LoginUser, constant.CreateCategory, constant.CreateTask, constant.ListTask,
	)
}
