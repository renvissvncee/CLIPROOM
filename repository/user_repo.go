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

func (r *UserRepository) NextUserId() uint {
	return r.storage.NextUserId()
}

// CREATE-функции

func (r *UserRepository) SaveUser(user *models.User) error {
	if user == nil {
		return errors.New("Ссылка на пользователя не найдена")
	}
	if user.ID <= 0 {
		return errors.New("Id пользователя не может быть <= 0")
	}
	r.storage.Users[user.ID] = *user
	return nil
}

// READ-функции

func (r *UserRepository) FindUserByLogin(login string) (models.User, error) {
	userMap := r.storage.Users
	for _, value := range userMap {
		if value.Login == login {
			return value, nil
		}
	}
	return models.User{}, ErrNotFound
}

func (r *UserRepository) FindUserById(id uint) (models.User, error) {
	if user, ok := r.storage.Users[id]; ok {
		return user, nil
	}
	return models.User{}, ErrNotFound
}

func (r *UserRepository) FindUserByEmail(email string) (models.User, error) {
	userMap := r.storage.Users
	for _, value := range userMap {
		if value.Email == email {
			return value, nil
		}
	}
	return models.User{}, ErrNotFound
}

func (r *UserRepository) FindAllUsers() ([]models.User, error) {
	usersSlice := make([]models.User, 0, len(r.storage.Users))
	for _, user := range r.storage.Users {
		usersSlice = append(usersSlice, user)
	}
	return usersSlice, nil
}

func (r *UserRepository) DeleteUser(id uint) error {
	userMap := r.storage.Users
	if _, ok := userMap[id]; !ok {
		return ErrNotFound
	}
	delete(userMap, id)
	return nil
}
