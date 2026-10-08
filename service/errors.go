package service

import "errors"

const noExcludeId uint = 0

var (
	ErrLoginTaken   = errors.New("логин уже занят")
	ErrEmailTaken   = errors.New("email уже занят")
	ErrUserNotFound = errors.New("пользователь не найден")
)
