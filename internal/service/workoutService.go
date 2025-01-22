package service

import (
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/repository"
)

func CreateWorkout(workout model.WorkoutPlan) (*model.WorkoutPlan, error) {
	return repository.CreateWorkout(workout)
}

func GetWorkout(id string) (*model.WorkoutPlan, error) {
	return repository.GetWorkout(id)
}

func UpdateWorkout(id string, workout model.WorkoutPlan) error {
	return repository.UpdateWorkout(id, workout)
}

func DeleteWorkout(id string) error {
	return repository.DeleteWorkout(id)
}
