package service

import (
    "github.com/rowjay007/walkit/internal/model"
    "github.com/rowjay007/walkit/internal/repository"
)

func LoginUser(login model.LoginRequest) (*model.LoginResponse, error) {
    return repository.LoginUser(login)
}

func RegisterUser(user model.User) error {
    return repository.RegisterUser(user)
}

func RequestPasswordReset(email string) error {
    return repository.RequestPasswordReset(email)
}

func ConfirmPasswordReset(request model.ConfirmPasswordResetRequest) error {
    return repository.ConfirmPasswordReset(request)
}