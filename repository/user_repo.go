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

func (usrepo *UserRepository) NextUserId() uint {
	return usrepo.storage.NextUserId()
}

// CREATE-функции

func (usrepo *UserRepository) SaveUser(user *models.User) error {
	if user == nil {
		return errors.New("Ссылка на пользователя не найдена")
	}
	if user.ID <= 0 {
		return errors.New("Id пользователя не может быть <= 0")
	}
	usrepo.storage.Users[user.ID] = *user
	return nil
}

// READ-функции

func (usrepo *UserRepository) FindUserByLogin(login string) (models.User, error) {
	userMap := usrepo.storage.Users
	for _, value := range userMap {
		if value.Login == login {
			return value, nil
		}
	}
	return models.User{}, ErrNotFound
}

func (usrepo *UserRepository) FindUserById(id uint) (models.User, error) {
	if user, ok := usrepo.storage.Users[id]; ok {
		return user, nil
	}
	return models.User{}, ErrNotFound
}

func (usrepo *UserRepository) FindUserByEmail(email string) (models.User, error) {
	userMap := usrepo.storage.Users
	for _, value := range userMap {
		if value.Email == email {
			return value, nil
		}
	}
	return models.User{}, ErrNotFound
}
