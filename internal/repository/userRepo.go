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

func UsersAPI() string {
	cfg := config.LoadConfig() 
	return cfg.BaseURL + "/collections/users/records"
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
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("registration failed: %s", string(body))
	}

	return nil
}

func GetUser(id string) (*model.User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", UsersAPI(), id))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get user: %s", string(body))
	}

	var user model.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &user, nil
}

func UpdateUser(id string, update model.UpdateUserRequest) error {
	updateJSON, err := json.Marshal(update)
	if err != nil {
		return fmt.Errorf("error marshaling update data: %w", err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", UsersAPI(), id), bytes.NewBuffer(updateJSON))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to update user: %s", string(body))
	}

	return nil
}

func DeleteUser(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", UsersAPI(), id), nil)
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete user: %s", string(body))
	}

	return nil
}
