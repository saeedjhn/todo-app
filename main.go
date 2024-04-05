package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type Task struct {
	Id         int
	UserId     int
	CategoryId int
	Title      string
	DueDate    string
	isDone     bool
}

type Category struct {
	Id     int
	UserId int
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
const UserStoragePath = "userStorageFile.txt"

func main() {
	fmt.Println("Hello to TODO app")

	loadUsersFromFile()

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

func loadUsersFromFile() {
	if _, err := os.Stat(UserStoragePath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		return
	}

	read, err := os.ReadFile(UserStoragePath)
	if err != nil {
		panic(err)
	}

	firstLetterToUpper := func(s string) string {

		if len(s) == 0 {
			return s
		}

		r := []rune(s)
		r[0] = unicode.ToUpper(r[0])

		return string(r)
	}

	for _, userSlice := range strings.Split(string(read), "\n") {
		if userSlice == "" {
			continue
		}

		values := strings.Split(userSlice, ", ")
		user := User{}
		// using for loop to iterate over the string
		for _, value := range values {
			parts := strings.Split(value, ":")
			if len(parts) != 2 {
				continue
			}
			key, val := parts[0], parts[1]
			v := reflect.ValueOf(&user).Elem()
			f := v.FieldByName(firstLetterToUpper(key))
			if !f.IsValid() {
				continue
			}
			if f.Type().Kind() == reflect.Int {
				age, err := strconv.Atoi(val)
				if err != nil {
					continue
				}
				f.SetInt(int64(age))
			} else {
				f.SetString(val)
			}
		}
		fmt.Printf("%+v\n", user)
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
		if task.UserId == authenticatedUser.Id {
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

	user := User{
		Id:       len(userStorage) + 1,
		Email:    email,
		Password: password,
	}

	userStorage = append(userStorage, user)

	f, err := os.OpenFile(UserStoragePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			fmt.Println("can`t close file,", err)
		}
	}(f)

	if err != nil {
		fmt.Println("can`t create or open file,", err)

		return
	}

	if _, err = f.WriteString(fmt.Sprintf(
		"id:%d, email:%s, password:%s\n",
		user.Id,
		user.Email,
		user.Password),
	); err != nil {
		fmt.Println("can`t write to file,", err)

		return
	}

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

	categoryId, err := strconv.Atoi(category)
	if err != nil {
		fmt.Printf("category id is not valid integer, %v", err)
		return
	}

	isFound := false
	for _, c := range categoryStorage {
		if c.Id == categoryId && c.UserId == authenticatedUser.Id {
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
		Id:         len(taskStorage) + 1,
		UserId:     authenticatedUser.Id,
		CategoryId: categoryId,
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
		Id:     len(categoryStorage) + 1,
		UserId: authenticatedUser.Id,
		Title:  title,
		Color:  color,
	})

	fmt.Printf("task is: %+v\n", categoryStorage[len(categoryStorage)-1])
}
