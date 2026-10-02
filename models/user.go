package models

type User struct {
	id       int
	login    string
	password string // пока без хэширования и обработки ошибок
	email    string
	name     string
	age      uint
}
