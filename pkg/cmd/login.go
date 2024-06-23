package cmd

import (
	"bufio"
	"fmt"
	"github.com/saeedjhn/todo-app/service/userservice"
	"log"
	"os"
)

func Login(us userservice.UserAdaptor) {
	fmt.Println("***** Login *****")

	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	users, err := us.Load()
	if err != nil {
		log.Fatalf("Can't load users %v", err)
	}

	for _, user := range users.Users {
		if user.Email == email && user.Password == password {
			fmt.Println("you`re logged in")
			AuthenticatedUser = &user

			break
		}
	}

	if AuthenticatedUser == nil {
		fmt.Println("the email or password is not correct")
	}
}
