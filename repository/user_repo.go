package repository

import (
	"cliproom/models"
	"cliproom/storage"
)

// Структура, отвечающая за запросы к "БД", связанные с пользователем
type UserRepository struct {
	storage *storage.Storage
}

func NewUserRepository(p_storage *storage.Storage) *UserRepository {
	return &UserRepository{storage: p_storage}
}

func (usrepo *UserRepository) AddUser(user *models.User) bool {
	if user == nil {
		return false
	}
	if user.Id <= 0 {
		return false
	}
	usrepo.storage.AddUser(user)
	return true
}

func (usrepo *UserRepository) NextId() uint {
	return usrepo.storage.GetMaxUserId() + 1
}
