package storage

import (
	"cliproom/models"
)

// Структура-замена БД, которая хранит основные данные (аналог подключения к БД)
type Storage struct {
	Users map[uint]models.User
	Clips map[uint]models.Clip
	// возможно, после, нужно будет поле UsersByLogin, где ключ - логин

	UserClips map[uint]map[uint]struct{}

	nextUserId uint
	nextClipId uint
}

func NewStorage() *Storage {
	return &Storage{
		Users:      map[uint]models.User{},
		Clips:      map[uint]models.Clip{},
		UserClips:  map[uint]uint{},
		nextUserId: 1,
		nextClipId: 1,
	}
}

func (p_s *Storage) NextUserId() uint {
	id := p_s.nextUserId
	p_s.nextUserId++
	return id
}

func (p_s *Storage) NextClipId() uint {
	id := p_s.nextClipId
	p_s.nextClipId++
	return id
}
