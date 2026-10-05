package service

import (
	"cliproom/models"
	"cliproom/repository"
	"errors"
	"strings"
)

// Структура, отвечающая за бизнес-лоигку, связанную с пользователем
type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(p_repositroy *repository.UserRepository) *UserService {
	return &UserService{userRepository: p_repositroy}
}

func (usserv *UserService) CreateUser(login, password, email, name string, age uint) (models.User, error) {
	// Валидация логина
	if login == "" {
		return models.User{}, errors.New("Логин не может быть пустым")
	} else if len(login) < 3 || len(login) > 30 {
		return models.User{}, errors.New("Логин должен быть верной длины: 3 <= login <= 30")
	} else if _, err := usserv.userRepository.FindUserByLogin(login); err == nil {
		return models.User{}, errors.New("Пользователь с таким логином уже существует")
	}
	// Валидация пароля
	if password == "" {
		return models.User{}, errors.New("Пароль не может быть пустым")
	} else if len(password) < 3 {
		return models.User{}, errors.New("Пароль не может быть меньше 3 символов")
	}
	// Валидация почты
	if !strings.Contains(email, "@") {
		return models.User{}, errors.New("Неверный формат эл. почты")
	}

	id := usserv.userRepository.NextUserId()
	newUser := models.User{
		Id:       id,
		Login:    login,
		Password: password,
		Email:    email,
		Name:     name,
		Age:      age,
	}
	return newUser, usserv.userRepository.AddUser(&newUser)
}

func (usserv *UserService) GetUserByLogin(login string) (models.User, error) {
	return usserv.userRepository.FindUserByLogin(login)
}

func (usserv *UserService) GetUserById(id uint) (models.User, error) {
	return usserv.userRepository.FindUserById(id)
}
