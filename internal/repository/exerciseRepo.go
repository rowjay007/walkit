package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/rowjay007/walkit/config"
	"github.com/rowjay007/walkit/internal/model"
)

func ExercisesAPI() string {
    return config.LoadConfig().BaseURL + "/collections/exercises/records" 
}


func CreateExercise(exercise model.Exercise) error {
	exerciseJSON, err := json.Marshal(exercise)
	if err != nil {
		return fmt.Errorf("error marshaling exercise data: %w", err)
	}

	resp, err := http.Post(ExercisesAPI(), "application/json", bytes.NewBuffer(exerciseJSON))
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to create exercise: %s", string(body))
	}
	return nil
}

func GetExercise(id string) (*model.Exercise, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", ExercisesAPI(), id))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get exercise: %s", string(body))
	}

	var exercise model.Exercise
	if err := json.NewDecoder(resp.Body).Decode(&exercise); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &exercise, nil
}

func ListExercises(filter model.ExerciseFilter) (*model.ExerciseListResponse, error) {
	query := url.Values{}

	if filter.Name != "" {
		query.Add("filter", fmt.Sprintf("name~'%s'", filter.Name))
	}
	if filter.Category != "" {
		query.Add("filter", fmt.Sprintf("category='%s'", filter.Category))
	}
	if filter.MuscleGroup != "" {
		query.Add("filter", fmt.Sprintf("muscleGroup~'%s'", filter.MuscleGroup))
	}
	if filter.SortBy != "" {
		sortPrefix := "+"
		if filter.SortOrder == "desc" {
			sortPrefix = "-"
		}
		query.Add("sort", fmt.Sprintf("%s%s", sortPrefix, filter.SortBy))
	}
	if filter.Page > 0 {
		query.Add("page", fmt.Sprintf("%d", filter.Page))
	}
	if filter.PerPage > 0 {
		query.Add("perPage", fmt.Sprintf("%d", filter.PerPage))
	}

	url := fmt.Sprintf("%s?%s", ExercisesAPI(), query.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	var response model.ExerciseListResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &response, nil
}

func UpdateExercise(id string, exercise model.Exercise) error {
	updateJSON, err := json.Marshal(exercise)
	if err != nil {
		return fmt.Errorf("error marshaling update data: %w", err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", ExercisesAPI(), id), bytes.NewBuffer(updateJSON))
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
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to update exercise: %s", string(body))
	}
	return nil
}

func DeleteExercise(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", ExercisesAPI(), id), nil)
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
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("failed to delete exercise: %s", string(body))
	}
	return nil
}
