package service

import (
    "github.com/rowjay007/walkit/internal/model"
    "github.com/rowjay007/walkit/internal/repository"
)

func GetUser(id string) (*model.User, error) {
    return repository.GetUser(id)
}

func ListUsers(filter model.UserFilter) (*model.UserListResponse, error) {
    return repository.ListUsers(filter)
}

func UpdateUser(id string, update model.UpdateUserRequest) error {
    return repository.UpdateUser(id, update)
}

func DeleteUser(id string) error {
    return repository.DeleteUser(id)
}
