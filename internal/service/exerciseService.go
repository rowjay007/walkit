package service

import "github.com/rowjay007/walkit/internal/model"
import "github.com/rowjay007/walkit/internal/repository"

// CreateExercise calls repository to create a new exercise.
func CreateExercise(exercise model.Exercise) error {
	return repository.CreateExercise(exercise)
}

// GetExercise fetches an exercise by its ID.
func GetExercise(id string) (*model.Exercise, error) {
	return repository.GetExercise(id)
}

// ListExercises lists exercises based on filter criteria.
func ListExercises(filter model.ExerciseFilter) (*model.ExerciseListResponse, error) {
	return repository.ListExercises(filter)
}

// UpdateExercise updates an existing exercise.
func UpdateExercise(id string, exercise model.Exercise) error {
	return repository.UpdateExercise(id, exercise)
}

// DeleteExercise deletes an exercise by its ID.
func DeleteExercise(id string) error {
	return repository.DeleteExercise(id)
}
