package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/saeedjhn/todo-app/constant"
	"github.com/saeedjhn/todo-app/pkg/alarm"
	"github.com/saeedjhn/todo-app/pkg/cmd"
	"github.com/saeedjhn/todo-app/service/taskservice"
	"github.com/saeedjhn/todo-app/service/userservice"
	"os"
)

func main() {
	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	fmt.Println("Hello to TODO app")
	fmt.Println(alarm.AllCommands())

	userSvc, taskSvc := cmd.ServiceSetup()

	for {
		runCommand(*command, userSvc, taskSvc)

		fmt.Println("please enter another command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(
	command string,
	us userservice.UserAdaptor,
	ts taskservice.TaskAdaptor,
) {
	if command != constant.RegisterUser && command != constant.Exit && cmd.AuthenticatedUser == nil {
		cmd.Login(us)

		return
	}

	switch command {
	case constant.CreateTask:
		cmd.CreateTask(ts)
	case constant.ListTask:
		cmd.ListTasks(ts)
	case constant.CreateCategory:
		cmd.CreateCategory()
	case constant.RegisterUser:
		cmd.Register(us)
	case constant.LoginUser:
		cmd.Login(us)
	case constant.Exit:
		os.Exit(0)
	default:
		fmt.Println(alarm.FieldCommand(command))
	}
}
