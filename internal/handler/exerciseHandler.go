package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"github.com/rowjay007/walkit/pkg/util"
)

// CreateExercise handler for creating a new exercise.
// CreateExercise godoc
// @Summary Create new exercise
// @Description Create a new exercise in the system
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param exercise body model.Exercise true "Exercise details"
// @Success 201 {object} model.Response{data=model.Exercise}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /exercises [post]
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

// GetExercise godoc
// @Summary Get exercise by ID
// @Description Get details of a specific exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Exercise ID"
// @Success 200 {object} model.Response{data=model.Exercise}
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /exercises/{id} [get]
func GetExercise(c *gin.Context) {
	id := c.Param("id")
	exercise, err := service.GetExercise(id)
	if err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(c, http.StatusOK, exercise)
}

// ListExercises godoc
// @Summary List exercises
// @Description Get a list of exercises with optional filtering
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param type query string false "Exercise type"
// @Success 200 {object} model.Response{data=[]model.Exercise}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Router /exercises [get]
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

// UpdateExercise godoc
// @Summary Update exercise
// @Description Update an existing exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Exercise ID"
// @Param exercise body model.Exercise true "Updated exercise details"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /exercises/{id} [put]
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

// DeleteExercise godoc
// @Summary Delete exercise
// @Description Delete an existing exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Exercise ID"
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /exercises/{id} [delete]
func DeleteExercise(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteExercise(id); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Exercise deleted successfully"})
}
