package service

import "errors"

var (
	ErrLoginTaken   = errors.New("логин уже занят")
	ErrEmailTaken   = errors.New("email уже занят")
	ErrUserNotFound = errors.New("пользователь не найден")
)
