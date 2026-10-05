package main

import (
	"cliproom/repository"
	"cliproom/service"
	"cliproom/storage"
	"fmt"
)

// Инициализируем основные слои: storage, repo, service
var db *storage.Storage = storage.NewStorage()
var userRepository *repository.UserRepository = repository.NewUserRepository(db)
var userService *service.UserService = service.NewUserService(userRepository)

func main() {
	userService.CreateUser(
		"renvissvnce",
		"123",
		"tyomafrutis@mail.ru",
		"Артем",
		20,
	)
	userService.CreateUser(
		"renvissvnce21",
		"123",
		"tyomafrutis@mail.ru",
		"Темыч",
		20,
	)

	fmt.Println(userService.GetUserById(1))
}
