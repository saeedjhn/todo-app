package filestore

import (
	"errors"
	"fmt"
	"github.com/saeedjhn/todo-app/domain"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type UserFileRepository struct {
	FilePath string
}

func NewUserFileRepository(filePath string) domain.UserRepository {
	return UserFileRepository{FilePath: filePath}
}

func (u UserFileRepository) Save(user domain.User) error {
	f, err := os.OpenFile(u.FilePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	defer func(f *os.File) error {
		err := f.Close()
		if err != nil {
			return fmt.Errorf("can`t close file, %s", err)
		}
		return nil
	}(f)

	if err != nil {
		return fmt.Errorf("can`t create or open file, %s", err)
	}

	if _, err = f.WriteString(fmt.Sprintf(
		"id:%d, email:%s, password:%s\n",
		user.Id,
		user.Email,
		user.Password),
	); err != nil {
		return fmt.Errorf("can`t write to file, %s", err)
	}

	return nil
}

func (u UserFileRepository) Load() ([]domain.User, error) {
	if _, err := os.Stat(u.FilePath); errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does not exist
		return nil, nil
	}

	read, err := os.ReadFile(u.FilePath)
	if err != nil {
		panic(err)
	}

	var users []domain.User
	for _, userSlice := range strings.Split(string(read), "\n") {
		if userSlice == "" {
			continue
		}

		values := strings.Split(userSlice, ", ")
		user := domain.User{}
		// using for loop to iterate over the string
		for _, value := range values {
			parts := strings.Split(value, ":")
			if len(parts) != 2 {
				continue
			}
			key, val := parts[0], parts[1]
			v := reflect.ValueOf(&user).Elem()
			f := v.FieldByName(u.firstLetterToUpper(key))
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
			users = append(users, user)
		}
	}
	return users, nil
}

func (u UserFileRepository) firstLetterToUpper(s string) string {
	if len(s) == 0 {
		return s
	}

	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])

	return string(r)
}
