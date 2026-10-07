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
	fmt.Println(userService.CreateUser(
		"renvissvnce",
		"123",
		"tyomafrutis@mail.ru",
		"Артем",
		20,
	))
	fmt.Println(userService.CreateUser(
		"renvissvnce21",
		"123",
		"tyomafrutis21@mail.ru",
		"Темыч",
		20,
	))
	newLogin := "renvissvnce34"
	newEmail := "tyomafrutis@mail.ru"
	fmt.Println(userService.UpdateUserInfo(2, &newLogin, &newEmail, nil, nil))
	fmt.Println(userService.GetUserById(2))
}
