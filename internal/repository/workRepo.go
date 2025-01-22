package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rowjay007/walkit/config"
	"github.com/rowjay007/walkit/internal/model"
)

func WorkoutAPI() string {
	return config.LoadConfig().BaseURL + "/collections/workouts/records"
}

func CreateWorkout(workout model.WorkoutPlan) (*model.WorkoutPlan, error) {
	workoutJSON, err := json.Marshal(workout)
	if err != nil {
		return nil, fmt.Errorf("error marshaling workout data: %w", err)
	}

	resp, err := http.Post(WorkoutAPI(), "application/json", bytes.NewBuffer(workoutJSON))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to create workout: %s", string(body))
	}

	var createdWorkout model.WorkoutPlan
	if err := json.NewDecoder(resp.Body).Decode(&createdWorkout); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}
	return &createdWorkout, nil
}

func GetWorkout(id string) (*model.WorkoutPlan, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", WorkoutAPI(), id))
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("failed to get workout: %s", string(body))
	}

	var workout model.WorkoutPlan
	if err := json.NewDecoder(resp.Body).Decode(&workout); err != nil {
		return nil, fmt.Errorf("error decoding response: %w", err)
	}

	return &workout, nil
}

func UpdateWorkout(id string, update model.WorkoutPlan) error {
	updateJSON, err := json.Marshal(update)
	if err != nil {
		return fmt.Errorf("error marshaling update data: %w", err)
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", WorkoutAPI(), id), bytes.NewBuffer(updateJSON))
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
		return fmt.Errorf("failed to update workout: %s", string(body))
	}

	return nil
}

func DeleteWorkout(id string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", WorkoutAPI(), id), nil)
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
		return fmt.Errorf("failed to delete workout: %s", string(body))
	}

	return nil
}
