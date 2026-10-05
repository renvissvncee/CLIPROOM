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
		return models.User{}, errors.New("Пользователь с данным логином уже существует")
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
	} else if _, err := usserv.userRepository.FindUserByEmail(email); err == nil {
		return models.User{}, errors.New("Пользователь с данной эл. почтой уже существует")
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

	err := usserv.userRepository.AddUser(&newUser)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func (usserv *UserService) GetUserByLogin(login string) (models.User, error) {
	if login == "" {
		return models.User{}, errors.New("Логин не может быть пустым")
	} else if len(login) < 3 || len(login) > 30 {
		return models.User{}, errors.New("Логин должен быть верной длины: 3 <= login <= 30")
	}
	return usserv.userRepository.FindUserByLogin(login)
}

func (usserv *UserService) GetUserById(id uint) (models.User, error) {
	if id == 0 {
		return models.User{}, errors.New("Id не может быть равен 0")
	}
	return usserv.userRepository.FindUserById(id)
}

func (usserv *UserService) GetUserByEmail(email string) (models.User, error) {
	if email == "" {
		return models.User{}, errors.New("Эл. почта не может быть пустым")
	} else if !strings.Contains(email, "@") {
		return models.User{}, errors.New("Неверный формат эл. почты")
	}
	return usserv.userRepository.FindUserByEmail(email)
}
