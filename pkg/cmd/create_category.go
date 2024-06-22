package cmd

import (
	"bufio"
	"fmt"
	"github.com/saeedjhn/todo-app/domain/entity"
	"os"
)

func CreateCategory() {
	fmt.Println("***** Create Category *****")
	scanner := bufio.NewScanner(os.Stdin)
	var title, color string

	fmt.Println("please enter the category title:")
	scanner.Scan()
	title = scanner.Text()

	fmt.Println("please enter the color category:")
	scanner.Scan()
	color = scanner.Text()

	CategoryStorage = append(CategoryStorage, entity.Category{
		Id:     len(CategoryStorage) + 1,
		UserId: AuthenticatedUser.Id,
		Title:  title,
		Color:  color,
	})

	fmt.Printf("task is: %+v\n", CategoryStorage[len(CategoryStorage)-1])
}
