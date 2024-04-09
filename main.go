package main

import (
	"bufio"
	"flag"
	"fmt"
	. "github.com/saeedjhn/todo-app/constant"
	. "github.com/saeedjhn/todo-app/domain"
	"github.com/saeedjhn/todo-app/repository/filestore"
	"github.com/saeedjhn/todo-app/repository/memorystore"
	"github.com/saeedjhn/todo-app/service"
	"os"
	"strconv"
)

var userStorage []User
var categoryStorage []Category

var authenticatedUser *User

func main() {
	userService := service.NewUserService(filestore.NewUserFileRepository(UserStoragePath))
	taskService := service.NewTaskService(memorystore.NewTaskMemoryRepository())

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	fmt.Println("Hello to TODO app")

	for {
		runCommand(*command, userService, taskService)

		fmt.Println("please enter another command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(
	command string,
	us UserAdaptorI,
	ts TaskAdaptorI,
) {
	if command != RegisterUser && command != Exit && authenticatedUser == nil {
		login()

		return
	}

	switch command {
	case CreateTask:
		createTask(ts)
	case ListTask:
		listTask(ts)
	case CreateCategory:
		createCategory()
	case RegisterUser:
		register(us)
	case LoginUser:
		login()
	case Exit:
		os.Exit(0)
	default:
		fmt.Println("command is not valid", command)
		fmt.Println(
			"You are authorized to the following commands:",
			CreateTask,
			CreateCategory,
			RegisterUser,
			LoginUser,
		)
	}
}

func register(us UserAdaptorI) {
	fmt.Println("***** Register *****")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email:")
	scanner.Scan()
	email = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email {
			fmt.Println("please enter try again email, email already exists:")
			scanner.Scan()
			email = scanner.Text()
		}
	}

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	u := User{
		Id:       len(userStorage) + 1,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, u)

	us.Save(u)

	fmt.Printf("user is: %+v\n", userStorage[len(userStorage)-1])
}

func login() {
	fmt.Println("***** Login *****")

	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range userStorage {
		if user.Email == email && user.Password == password {
			fmt.Println("you`re logged in")
			authenticatedUser = &user

			break
		}
	}

	if authenticatedUser == nil {
		fmt.Println("the email or password is not correct")
	}
}

func createTask(ts TaskAdaptorI) {
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

	task, err := ts.Create(TaskCreateRequest{
		AuthenticatedUserId: authenticatedUser.Id,
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

func listTask(ts TaskAdaptorI) {
	tasks, err := ts.List(TaskListRequest{AuthenticatedUserId: authenticatedUser.Id})
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println("User tasks:", tasks.Tasks)
}

func createCategory() {
	fmt.Println("***** Create Category *****")
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the color category:")
	scanner.Scan()
	color = scanner.Text()

	categoryStorage = append(categoryStorage, Category{
		Id:     len(categoryStorage) + 1,
		UserId: authenticatedUser.Id,
		Title:  title,
		Color:  color,
	})

	fmt.Printf("task is: %+v\n", categoryStorage[len(categoryStorage)-1])
}
