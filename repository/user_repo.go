package repository

import (
	"cliproom/models"
	"cliproom/storage"
	"errors"
)

// Структура, отвечающая за запросы к "БД", связанные с пользователем
type UserRepository struct {
	storage *storage.Storage
}

func NewUserRepository(p_storage *storage.Storage) *UserRepository {
	return &UserRepository{storage: p_storage}
}

func (usrepo *UserRepository) AddUser(user *models.User) error {
	if user == nil {
		return errors.New("Ссылка на пользователя не найдена")
	}
	if user.Id <= 0 {
		return errors.New("Id пользователя не может быть <= 0")
	}
	usrepo.storage.Users[user.Id] = *user
	return nil
}

func (usrepo *UserRepository) FindUserByLogin(login string) (models.User, error) {
	userMap := usrepo.storage.Users
	for _, value := range userMap {
		if value.Login == login {
			return value, nil
		}
	}
	return models.User{}, errors.New("Пользователь не найден")
}

func (usrepo *UserRepository) FindUserById(id uint) (models.User, error) {
	if user, ok := usrepo.storage.Users[id]; ok {
		return user, nil
	}
	return models.User{}, errors.New("Пользователь не найден")
}

func (usrepo *UserRepository) NextUserId() uint {
	return usrepo.storage.NextUserId()
}
