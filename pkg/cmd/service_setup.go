package cmd

import (
	"github.com/saeedjhn/todo-app/constant"
	"github.com/saeedjhn/todo-app/repository/filestore"
	"github.com/saeedjhn/todo-app/repository/memorystore"
	"github.com/saeedjhn/todo-app/service/taskservice"
	"github.com/saeedjhn/todo-app/service/userservice"
)

func ServiceSetup() (userservice.UserAdaptor, taskservice.TaskAdaptor) {
	userSvc := userservice.New(filestore.New(constant.UserStoragePath))
	taskSvc := taskservice.New(memorystore.New())

	return userSvc, taskSvc
}
