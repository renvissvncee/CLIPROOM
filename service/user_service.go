package service

import (
	"cliproom/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(p_repositroy *repository.UserRepository) *UserService {
	return &UserService{repository: p_repositroy}
}
