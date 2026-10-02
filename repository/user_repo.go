package repository

import "cliproom/storage"

type UserRepository struct {
	storage *storage.Storage
}

func NewUserRepository(p_storage *storage.Storage) *UserRepository {
	return &UserRepository{storage: p_storage}
}
