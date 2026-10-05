package service

import (
	"cliproom/models"
	"cliproom/repository"
	"strings"
)

// Структура, отвечающая за бизнес-лоигку, связанную с пользователем
type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(p_repositroy *repository.UserRepository) *UserService {
	return &UserService{userRepository: p_repositroy}
}

func (usserv UserService) CreateUser(login, password, email, name string, age uint) (models.User, bool) {
	// Валидация
	if login == "" {
		return models.User{}, false
	} else if password == "" {
		return models.User{}, false
	} else if email == "" || strings.Index(email, "@") == -1 {
		return models.User{}, false
	} else if name == "" {
		return models.User{}, false
	} else if age < 0 {
		return models.User{}, false
	}

	if

	id := usserv.userRepository.NextId()
	newUser := models.User{
		Id:       id,
		Login:    login,
		Password: password,
		Email:    email,
		Name:     name,
		Age:      age,
	}

}
