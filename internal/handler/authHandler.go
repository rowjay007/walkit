package handler

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/model"
	"github.com/rowjay007/walkit/internal/service"
	"github.com/rowjay007/walkit/pkg/util"
)

// LoginUser godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param login body model.LoginRequest true "Login credentials"
// @Success 200 {object} model.Response{data=model.LoginResponse}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Router /auth/login [post]
func LoginUser(c *gin.Context) {
    var login model.LoginRequest
    if err := c.ShouldBindJSON(&login); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{Error: err.Error()})
        return
    }

    if err := validateLoginInput(login); err != nil {
        util.RespondWithError(c, http.StatusBadRequest, err.Error())
        return
    }

    response, err := service.LoginUser(login)
    if err != nil {

        
        statusCode := determineStatusCode(err)
        util.RespondWithError(c, statusCode, "Login failed. Please check your credentials.")
        return
    }

    util.RespondWithJSON(c, http.StatusOK, response)
}

func validateLoginInput(login model.LoginRequest) error {
    if strings.TrimSpace(login.Identity) == "" {
        return fmt.Errorf("identity cannot be empty")
    }
    if len(login.Password) < 8 {
        return fmt.Errorf("password must be at least 8 characters")
    }
    return nil
}

func determineStatusCode(err error) int {
    if strings.Contains(strings.ToLower(err.Error()), "invalid credentials") {
        return http.StatusUnauthorized
    }
    if strings.Contains(strings.ToLower(err.Error()), "not found") {
        return http.StatusNotFound
    }
    return http.StatusInternalServerError
}

// RegisterUser godoc
// @Summary Register new user
// @Description Create a new user account
// @Tags auth
// @Accept json
// @Produce json
// @Param user body model.User true "User registration details"
// @Success 201 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/register [post]
func RegisterUser(c *gin.Context) {
    var user model.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{Error: err.Error()})
        return
    }

    if err := service.RegisterUser(user); err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    util.RespondWithJSON(c, http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// RequestPasswordReset godoc
// @Summary Request password reset
// @Description Send a password reset email to the user
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.PasswordResetRequest true "Password reset request"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/password-reset [post]
func RequestPasswordReset(c *gin.Context) {
    var request model.PasswordResetRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{Error: "Invalid request payload"})
        return
    }

    if err := service.RequestPasswordReset(request.Email); err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, model.Response{Message: "Password reset email sent"})
}

// ConfirmPasswordReset godoc
// @Summary Confirm password reset
// @Description Reset user's password using the reset token
// @Tags auth
// @Accept json
// @Produce json
// @Param request body model.ConfirmPasswordResetRequest true "Password reset confirmation"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /auth/password-reset/confirm [post]
func ConfirmPasswordReset(c *gin.Context) {
    var request model.ConfirmPasswordResetRequest
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, model.Response{Error: "Invalid request payload"})
        return
    }

    if err := service.ConfirmPasswordReset(request); err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }

    c.JSON(http.StatusOK, model.Response{Message: "Password reset successfully"})
}



