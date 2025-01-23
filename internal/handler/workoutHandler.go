package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
)


// CreateWorkout godoc
// @Summary Create new workout
// @Description Create a new workout plan
// @Tags workouts
// @Accept json
// @Produce json
// @Param workout body model.WorkoutPlan true "Workout details"
// @Success 201 {object} model.Response{data=model.WorkoutPlan}
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /workouts [post]
func CreateWorkout(c *gin.Context) {
    var workout model.WorkoutPlan
    if err := c.ShouldBindJSON(&workout); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{Error: err.Error()})
        return
    }

    createdWorkout, err := service.CreateWorkout(workout)
    if err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusCreated, model.Response{Data: createdWorkout})
}

// GetWorkout godoc
// @Summary Get workout by ID
// @Description Get details of a specific workout
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path string true "Workout ID"
// @Success 200 {object} model.Response{data=model.WorkoutPlan}
// @Failure 404 {object} model.Response
// @Router /workouts/{id} [get]
func GetWorkout(c *gin.Context) {
    id := c.Param("id")
    workout, err := service.GetWorkout(id)
    if err != nil {
        c.JSON(http.StatusNotFound, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, model.Response{Data: workout})
}

// UpdateWorkout godoc
// @Summary Update workout
// @Description Update an existing workout plan
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path string true "Workout ID"
// @Param workout body model.WorkoutPlan true "Updated workout details"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /workouts/{id} [put]
func UpdateWorkout(c *gin.Context) {
    id := c.Param("id")
    var workout model.WorkoutPlan
    if err := c.ShouldBindJSON(&workout); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{Error: err.Error()})
        return
    }

    if err := service.UpdateWorkout(id, workout); err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, model.Response{Message: "workout updated successfully"})
}

// DeleteWorkout godoc
// @Summary Delete workout
// @Description Delete an existing workout plan
// @Tags workouts
// @Accept json
// @Produce json
// @Param id path string true "Workout ID"
// @Success 200 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /workouts/{id} [delete]
func DeleteWorkout(c *gin.Context) {
    id := c.Param("id")
    if err := service.DeleteWorkout(id); err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, model.Response{Message: "workout deleted successfully"})
}
