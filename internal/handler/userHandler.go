package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"github.com/rowjay007/walkit/pkg/util"
)

// LoginUser handles user login.
func LoginUser(c *gin.Context) {
	var login model.LoginRequest
	if err := c.ShouldBindJSON(&login); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	response, err := service.LoginUser(login)
	if err != nil {
		util.RespondWithError(c, http.StatusUnauthorized, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, response)
}

// RequestPasswordReset handles password reset requests.
func RequestPasswordReset(c *gin.Context) {
	var request model.PasswordResetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := service.RequestPasswordReset(request.Email); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Password reset email sent"})
}

// ConfirmPasswordReset handles password reset confirmation.
func ConfirmPasswordReset(c *gin.Context) {
	var request model.ConfirmPasswordResetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := service.ConfirmPasswordReset(request); err != nil {
		util.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Password reset successfully"})
}
