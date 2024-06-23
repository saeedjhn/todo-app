package cmd

import (
	"bufio"
	"fmt"
	"github.com/saeedjhn/todo-app/domain/dto/userdto"
	"github.com/saeedjhn/todo-app/service/userservice"
	"log"
	"os"
)

func Register(us userservice.UserAdaptor) {
	fmt.Println("***** Register *****")
	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email:")
	scanner.Scan()
	email = scanner.Text()

	users, err := us.Load()
	if err != nil {
		log.Fatalf("Can't load users %v", err)
	}

	for _, user := range users.Users {
		if user.Email == email {
			fmt.Println("please enter try again email, email already exists:")
			scanner.Scan()
			email = scanner.Text()
		}
	}

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	us.Save(userdto.SaveRequest{
		Email:    email,
		Password: password,
	})

	fmt.Println("user is register successfully")
}
