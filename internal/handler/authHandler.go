package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"github.com/rowjay007/walkit/pkg/util"
)

// RegisterUser handles user registration.
func RegisterUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := service.RegisterUser(user); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// GetUser retrieves a user by ID.
func GetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := service.GetUser(id)
	if err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, user)
}

// UpdateUser updates user information.
func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var update model.UpdateUserRequest
	if err := c.ShouldBindJSON(&update); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := service.UpdateUser(id, update); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User updated successfully"})
}

// DeleteUser deletes a user by ID.
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteUser(id); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}
