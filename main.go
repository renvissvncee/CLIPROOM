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

var p *int

func main() {
	fmt.Println(p)
}
