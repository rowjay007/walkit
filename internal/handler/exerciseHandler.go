package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"net/http"
)

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
