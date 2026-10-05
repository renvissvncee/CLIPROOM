package storage

import "cliproom/models"

// Структура-замена БД, которая хранит основные данные (аналог подключения к БД)
type Storage struct {
	Users []models.User
	Clips []models.Clip

	UserClips map[uint]uint

	maxUserId uint
	maxClipId uint
}

func NewStorage() *Storage {
	return &Storage{
		Users:     []models.User{},
		Clips:     []models.Clip{},
		UserClips: map[uint]uint{},
		maxUserId: 0,
		maxClipId: 0,
	}
}

func (p_s *Storage) AddUser(user *models.User) bool {
	p_s.Users = append(p_s.Users, *user)
	return true
}

func (p_s *Storage) GetMaxUserId() uint {
	return p_s.maxUserId
}
