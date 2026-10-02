package storage

import "cliproom/models"

type Storage struct {
	Users []models.User
	Clips []models.Clip

	UserClips map[int]int

	maxUserId int
	maxClipId int
}

func NewStorage() *Storage {
	return &Storage{
		Users:     []models.User{},
		Clips:     []models.Clip{},
		UserClips: map[int]int{},
		maxUserId: 0,
		maxClipId: 0,
	}
}

func (s Storage) addUser(user *models.User) {
	if user == nil {
		return
	}

}
