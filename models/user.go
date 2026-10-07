package models

type User struct {
	ID       uint
	Login    string
	Password string // пока без хэширования и обработки ошибок
	Email    string
	Name     string
	Age      uint
}
