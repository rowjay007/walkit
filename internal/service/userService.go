package service

import (
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/repository"
)

func RegisterUser(user model.User) error {
	return repository.RegisterUser(user)
}

func GetUser(id string) (*model.User, error) {
	return repository.GetUser(id)
}

func UpdateUser(id string, update model.UpdateUserRequest) error {
	return repository.UpdateUser(id, update)
}

func DeleteUser(id string) error {
	return repository.DeleteUser(id)
}
