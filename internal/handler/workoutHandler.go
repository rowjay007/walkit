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


func GetWorkout(c *gin.Context) {
    id := c.Param("id")
    workout, err := service.GetWorkout(id)
    if err != nil {
        c.JSON(http.StatusNotFound, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, model.Response{Data: workout})
}


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

func DeleteWorkout(c *gin.Context) {
    id := c.Param("id")
    if err := service.DeleteWorkout(id); err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, model.Response{Message: "workout deleted successfully"})
}
