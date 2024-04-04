package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
)

type User struct {
	ID       int
	Email    string
	Password string
}

type Task struct {
	ID         int
	UserID     int
	CategoryID int
	Title      string
	DueDate    string
	isDone     bool
}

type Category struct {
	ID     int
	UserID int
	Title  string
	Color  string
}

var userStorage []User
var categoryStorage []Category
var taskStorage []Task

var authenticatedUser *User

const (
	CreateTask     = "create-task"
	ListTask       = "list-task"
	CreateCategory = "create-category"
	RegisterUser   = "register-user"
	LoginUser      = "login-user"
	Exit           = "exit"
)

func main() {
	fmt.Println("Hello to TODO app")

	command := flag.String("command", "no-command", "command to run")
	flag.Parse()

	for {
		runCommand(*command)

		fmt.Println("please enter another command:")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		*command = scanner.Text()
	}
}

func runCommand(command string) {
	if command != RegisterUser && command != Exit && authenticatedUser == nil {
		login()

		return
	}

	switch command {
	case CreateTask:
		createTask()
	case ListTask:
		listTask()
	case CreateCategory:
		createCategory()
	case RegisterUser:
		register()
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

func listTask() {
	for _, task := range taskStorage {
		if task.UserID == authenticatedUser.ID {
			fmt.Printf("%+v\n", task)
		}
	}
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

func register() {
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

	userStorage = append(userStorage, User{
		ID:       len(userStorage) + 1,
		Email:    email,
		Password: password,
	})

	fmt.Printf("user is: %+v\n", userStorage[len(userStorage)-1])
}

func createTask() {
	fmt.Println("***** Create Task ******")
	var title, category, duedate string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("please enter the task title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the task category:")
	scanner.Scan()
	category = scanner.Text()

	categoryID, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category id is not valid integer, %v", err)
		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.ID == categoryID && c.UserID == authenticatedUser.ID {
			isFound = true

			break
		}
	}

	if !isFound {
		fmt.Println("category-id is not found")

		return
	}

	fmt.Println("please enter the task due date:")
	scanner.Scan()
	duedate = scanner.Text()

	taskStorage = append(taskStorage, Task{
		ID:         len(taskStorage) + 1,
		UserID:     authenticatedUser.ID,
		CategoryID: categoryID,
		Title:      title,
		DueDate:    duedate,
		isDone:     false,
	})

	fmt.Printf("task is: %+v\n", taskStorage[len(taskStorage)-1])
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
		ID:     len(categoryStorage) + 1,
		UserID: authenticatedUser.ID,
		Title:  title,
		Color:  color,
	})

	fmt.Printf("task is: %+v\n", categoryStorage[len(categoryStorage)-1])
}
