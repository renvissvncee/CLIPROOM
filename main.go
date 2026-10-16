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

var clipRepository *repository.ClipRepository = repository.NewClipRepository(db)
var clipService *service.ClipService = service.NewClipService(clipRepository)

func main() {
	fmt.Println(clipService.CreateClip(
		"POWER",
		"Kanye West",
		"",
		"abc.ru",
		[]string{"hip-hop"},
		[]string{"energetic"},
	))

	fmt.Println(clipService.CreateClip(
		"Homecoming",
		"Kanye West",
		"",
		"abc.ru",
		[]string{"hip-hop"},
		[]string{"energetic"},
	))
}
