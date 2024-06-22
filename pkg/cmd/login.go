package cmd

import (
	"bufio"
	"fmt"
	"os"
)

func Login() {
	fmt.Println("***** Login *****")

	scanner := bufio.NewScanner(os.Stdin)
	var email, password string

	fmt.Println("please enter the email:")
	scanner.Scan()
	email = scanner.Text()

	fmt.Println("please enter the password:")
	scanner.Scan()
	password = scanner.Text()

	for _, user := range UserStorage {
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
