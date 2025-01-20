package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"github.com/rowjay007/walkit/pkg/util"
)

// CreateExercise handler for creating a new exercise.
func CreateExercise(c *gin.Context) {
	var exercise model.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := service.CreateExercise(exercise); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusCreated, gin.H{"message": "Exercise created successfully"})
}

// GetExercise handler to get a single exercise by ID.
func GetExercise(c *gin.Context) {
	id := c.Param("id")
	exercise, err := service.GetExercise(id)
	if err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(c, http.StatusOK, exercise)
}

// ListExercises handler to list exercises with optional filters.
func ListExercises(c *gin.Context) {
	var filter model.ExerciseFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid query parameters")
		return
	}

	response, err := service.ListExercises(filter)
	if err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, response)
}

// UpdateExercise handler to update an exercise by ID.
func UpdateExercise(c *gin.Context) {
	id := c.Param("id")
	var exercise model.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := service.UpdateExercise(id, exercise); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Exercise updated successfully"})
}

// DeleteExercise handler to delete an exercise by ID.
func DeleteExercise(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteExercise(id); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Exercise deleted successfully"})
}
