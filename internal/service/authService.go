package service

import (
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/repository"
)

func LoginUser(login model.LoginRequest) (*model.LoginResponse, error) {
	return repository.LoginUser(login)
}

func RequestPasswordReset(email string) error {

	return nil
}

func ConfirmPasswordReset(request model.ConfirmPasswordResetRequest) error {
	return nil
}

