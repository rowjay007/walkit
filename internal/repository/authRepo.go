package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/config"
)

func AuthAPI() string {
	return config.LoadConfig().BaseURL + "/collections/users/auth-with-password"
}

func LoginUser(login model.LoginRequest) (*model.LoginResponse, error) {
	loginJSON, err := json.Marshal(login)
	if err != nil {
		return nil, fmt.Errorf("error marshaling login data: %w", err)
	}

	resp, err := http.Post(AuthAPI(), "application/json", bytes.NewBuffer(loginJSON))
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

func RequestPasswordReset(email string) error {
	resetRequest := model.PasswordResetRequest{Email: email}
	resetJSON, err := json.Marshal(resetRequest)
	if err != nil {
		return fmt.Errorf("error marshaling reset request: %w", err)
	}

	resp, err := http.Post(config.LoadConfig().BaseURL+"/collections/users/request-password-reset", "application/json", bytes.NewBuffer(resetJSON))
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

func ConfirmPasswordReset(request model.ConfirmPasswordResetRequest) error {
	resetJSON, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("error marshaling reset confirmation: %w", err)
	}

	resp, err := http.Post(config.LoadConfig().BaseURL+"/collections/users/confirm-password-reset", "application/json", bytes.NewBuffer(resetJSON))
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
