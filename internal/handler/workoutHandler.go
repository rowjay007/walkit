package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
)

func CreateWorkout(c *gin.Context) {
	var workout model.WorkoutPlan
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdWorkout, err := service.CreateWorkout(workout)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdWorkout)
}

func GetWorkout(c *gin.Context) {
	id := c.Param("id")
	workout, err := service.GetWorkout(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, workout)
}

func UpdateWorkout(c *gin.Context) {
	id := c.Param("id")
	var workout model.WorkoutPlan
	if err := c.ShouldBindJSON(&workout); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.UpdateWorkout(id, workout); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "workout updated successfully"})
}

func DeleteWorkout(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteWorkout(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "workout deleted successfully"})
}
