package repository

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
    "github.com/rowjay007/walkit/internal/model"
    "github.com/rowjay007/walkit/config"
)

func UsersAPI() string {
    return config.LoadConfig().BaseURL + "/collections/users/records"
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

func ListUsers(filter model.UserFilter) (*model.UserListResponse, error) {
    query := url.Values{}
    
    // Build filter conditions
    if filter.Username != "" {
        query.Add("filter", fmt.Sprintf("username~'%s'", filter.Username))
    }
    if filter.Email != "" {
        query.Add("filter", fmt.Sprintf("email~'%s'", filter.Email))
    }
    if filter.FitnessGoal != "" {
        query.Add("filter", fmt.Sprintf("fitnessGoal='%s'", filter.FitnessGoal))
    }
    if filter.ActivityLevel != "" {
        query.Add("filter", fmt.Sprintf("activityLevel='%s'", filter.ActivityLevel))
    }

    // Add sorting
    if filter.SortBy != "" {
        sortPrefix := "+"
        if filter.SortOrder == "desc" {
            sortPrefix = "-"
        }
        query.Add("sort", fmt.Sprintf("%s%s", sortPrefix, filter.SortBy))
    }

    // Add pagination
    if filter.Page > 0 {
        query.Add("page", fmt.Sprintf("%d", filter.Page))
    }
    if filter.PerPage > 0 {
        query.Add("perPage", fmt.Sprintf("%d", filter.PerPage))
    }

    // Make request
    url := fmt.Sprintf("%s?%s", UsersAPI(), query.Encode())
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("error making request: %w", err)
    }
    defer resp.Body.Close()

    var response model.UserListResponse
    if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
        return nil, fmt.Errorf("error decoding response: %w", err)
    }

    return &response, nil
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

		if resp.StatusCode != http.StatusOK {
			body, _ := ioutil.ReadAll(resp.Body)
			return fmt.Errorf("failed to delete user: %s", string(body))
		}

		return nil
	}