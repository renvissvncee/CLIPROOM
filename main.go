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
		"POWER",
		"Drake",
		"",
		"abc.ru",
		[]string{"hip-hop"},
		[]string{"energetic"},
	))

	fmt.Println(clipService.GetAllClips())
	newVar := "jdgsjdjrg.ru"
	var newTagsRaw []string = []string{"rock", "jazz"}
	var newMoodsRaw []string = []string{"dark", "sad"}
	fmt.Println(clipService.UpdateClipInfo(1, nil, nil, nil, &newVar, newTagsRaw, newMoodsRaw))
}
