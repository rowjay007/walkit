package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/rowjay007/walkit/config"
	"github.com/rowjay007/walkit/internal/model"
)

func AuthAPI() string {
    return config.LoadConfig().BaseURL + "/collections/users"
}


func RegisterUser(user model.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("error marshaling user data: %w", err)
	}

	resp, err := http.Post(UsersAPI(), "application/json", bytes.NewBuffer(userJSON))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("registration failed: %s", string(body))
	}

	return nil
}

func LoginUser(login model.LoginRequest) (*model.LoginResponse, error) {
    if login.Identity == "" || login.Password == "" {
        return nil, fmt.Errorf("identity and password are required")
    }

    loginJSON, err := json.Marshal(map[string]interface{}{
        "identity": login.Identity,
        "password": login.Password,
    })
    if err != nil {
        return nil, fmt.Errorf("error marshaling login data: %w", err)
    }

    req, err := http.NewRequest("POST", AuthAPI()+"/auth-with-password", bytes.NewBuffer(loginJSON))
    if err != nil {
        return nil, fmt.Errorf("error creating request: %w", err)
    }
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{
        Timeout: 10 * time.Second,
    }
    
    resp, err := client.Do(req)
    if err != nil {
        return nil, fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, fmt.Errorf("error reading response body: %w", err)
    }

    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("login failed: %s", string(body))
    }

    var response struct {
        Token string `json:"token"`
        Record struct {
            ID              string    `json:"id"`
            Username        string    `json:"username"`
            Email           string    `json:"email"`
            EmailVisibility bool      `json:"emailVisibility"`
            Created         string    `json:"created"`
            Updated        string    `json:"updated"`
            Verified       bool      `json:"verified"`
            FitnessGoal    string    `json:"fitnessGoal"`
            ActivityLevel  string    `json:"activityLevel"`
            Avatar         string    `json:"avatar"`
        } `json:"record"`
    }

    if err := json.Unmarshal(body, &response); err != nil {
        return nil, fmt.Errorf("error decoding response: %w", err)
    }

    loginResponse := &model.LoginResponse{
        Token: response.Token,
        User: model.User{
            ID:              response.Record.ID,
            Username:        response.Record.Username,
            Email:           response.Record.Email,
            EmailVisibility: response.Record.EmailVisibility,
            FitnessGoal:    response.Record.FitnessGoal,
            ActivityLevel:   response.Record.ActivityLevel,
            Avatar:         response.Record.Avatar,
            CreatedAt:      response.Record.Created,
            UpdatedAt:      response.Record.Updated,
        },
    }

    return loginResponse, nil
}

func RequestPasswordReset(email string) error {
    resetRequest := model.PasswordResetRequest{Email: email}
    resetJSON, err := json.Marshal(resetRequest)
    if err != nil {
        return fmt.Errorf("error marshaling reset request: %w", err)
    }

    resp, err := http.Post(AuthAPI()+"/request-password-reset", "application/json", bytes.NewBuffer(resetJSON))
    if err != nil {
        return fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("password reset request failed: %s", string(body))
    }

    return nil
}

func ConfirmPasswordReset(request model.ConfirmPasswordResetRequest) error {
    resetJSON, err := json.Marshal(request)
    if err != nil {
        return fmt.Errorf("error marshaling reset confirmation: %w", err)
    }

    resp, err := http.Post(AuthAPI()+"/confirm-password-reset", "application/json", bytes.NewBuffer(resetJSON))
    if err != nil {
        return fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("password reset confirmation failed: %s", string(body))
    }

    return nil
}