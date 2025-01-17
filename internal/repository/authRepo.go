package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rowjay007/walkit/config"
	"github.com/rowjay007/walkit/internal/model"
)

// LoginUser handles user login by sending a request to the authentication API.
func LoginUser(login model.LoginRequest) (*model.LoginResponse, error) {
	loginJSON, err := json.Marshal(login)
	if err != nil {
		return nil, fmt.Errorf("error marshaling login data: %w", err)
	}

	resp, err := http.Post(config.AuthAPI, "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("login failed: %s", string(body))
	}

	var loginResponse model.LoginResponse
	if err := json.NewDecoder(resp.Body).Decode(&loginResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &loginResponse, nil
}

// RequestPasswordReset handles the password reset request.
func RequestPasswordReset(email string) error {
	resetRequest := model.PasswordResetRequest{Email: email}
	resetJSON, err := json.Marshal(resetRequest)
	if err != nil {
		return fmt.Errorf("error marshaling reset request: %w", err)
	}

	resp, err := http.Post(config.ResetAPI, "application/json", bytes.NewBuffer(resetJSON))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("password reset request failed: %s", string(body))
	}

	return nil
}

// ConfirmPasswordReset handles the confirmation of a password reset.
func ConfirmPasswordReset(request model.ConfirmPasswordResetRequest) error {
	resetJSON, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling reset confirmation: %w", err)
	}

	resp, err := http.Post(config.ConfirmResetAPI, "application/json", bytes.NewBuffer(resetJSON))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("password reset confirmation failed: %s", string(body))
	}

	return nil
}
