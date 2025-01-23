package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"net/http"
)

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
		c.JSON(http.StatusBadRequest, model.Response{Error: "Invalid request payload"})
		return
	}

	if err := service.CreateExercise(exercise); err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, model.Response{Message: "Exercise created successfully"})
}

// @Summary Get exercise by ID
// @Description Get details of a specific exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Exercise ID"
// @Success 200 {object} model.Response{data=model.Exercise}
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /exercises/{id} [get]
func GetExercise(c *gin.Context) {
	id := c.Param("id")
	exercise, err := service.GetExercise(id)
	if err != nil {
		if err.Error() == "exercise not found" {
			c.JSON(http.StatusNotFound, model.Response{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.Response{Data: exercise})
}

// @Summary List exercises
// @Description Get a list of exercises with optional filtering
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} model.Response{data=[]model.Exercise}
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /exercises [get]
func ListExercises(c *gin.Context) {
	var filter model.ExerciseFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Error: "Invalid query parameters"})
		return
	}

	response, err := service.ListExercises(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{
		Data:    response,
		Message: "Exercises retrieved successfully",
	})
}

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
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /exercises/{id} [put]
func UpdateExercise(c *gin.Context) {
	id := c.Param("id")
	var exercise model.Exercise
	if err := c.ShouldBindJSON(&exercise); err != nil {
		c.JSON(http.StatusBadRequest, model.Response{Error: "Invalid request payload"})
		return
	}

	if err := service.UpdateExercise(id, exercise); err != nil {
		if err.Error() == "exercise not found" {
			c.JSON(http.StatusNotFound, model.Response{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, model.Response{Message: "Exercise updated successfully"})
}

// @Summary Delete exercise
// @Description Delete an existing exercise
// @Tags exercises
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "Exercise ID"
// @Success 200 {object} model.Response
// @Failure 404 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /exercises/{id} [delete]
func DeleteExercise(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteExercise(id); err != nil {
		if err.Error() == "exercise not found" {
			c.JSON(http.StatusNotFound, model.Response{Error: err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, model.Response{Message: "Exercise deleted successfully"})
}
