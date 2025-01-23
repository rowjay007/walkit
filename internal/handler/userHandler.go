package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/rowjay007/walkit/internal/model"
    "github.com/rowjay007/walkit/internal/service"
    "github.com/rowjay007/walkit/pkg/util"
)

// @Summary Get current user
// @Description Get the profile of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response{data=model.User}
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /users/me [get]
func GetMe(c *gin.Context) {
    userID := c.GetString("userID")
    user, err := service.GetUser(userID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, model.Response{Error: err.Error()})
        return
    }
    c.JSON(http.StatusOK, model.Response{Data: user})
}

// @Summary Update current user
// @Description Update the profile of the currently authenticated user
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param user body model.UpdateUserRequest true "User update details"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Router /users/me [put]
func UpdateMe(c *gin.Context) {
    userID := c.GetString("userID")
    var update model.UpdateUserRequest
    if err := c.ShouldBindJSON(&update); err != nil {
        util.RespondWithError(c, http.StatusBadRequest, "Invalid request payload")
        return
    }

    if err := service.UpdateUser(userID, update); err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Profile updated successfully"})
}

// @Summary Delete current user
// @Description Delete the currently authenticated user's account
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 500 {object} model.Response
// @Router /users/me [delete]
func DeleteMe(c *gin.Context) {
    userID := c.GetString("userID")
    if err := service.DeleteUser(userID); err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

// @Summary List users
// @Description Get a list of users with optional filtering
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} model.Response{data=[]model.User}
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Router /users [get]
func ListUsers(c *gin.Context) {
    var filter model.UserFilter
    if err := c.ShouldBindQuery(&filter); err != nil {
        util.RespondWithError(c, http.StatusBadRequest, "Invalid query parameters")
        return
    }

    response, err := service.ListUsers(filter)
    if err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }

    util.RespondWithJSON(c, http.StatusOK, response)
}

// @Summary Get user by ID
// @Description Get details of a specific user
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "User ID"
// @Success 200 {object} model.Response{data=model.User}
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := service.GetUser(id)
    if err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, user)
}

// @Summary Update user
// @Description Update an existing user
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "User ID"
// @Param user body model.UpdateUserRequest true "Updated user details"
// @Success 200 {object} model.Response
// @Failure 400 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /users/{id} [put]
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

// @Summary Delete user
// @Description Delete an existing user
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path string true "User ID"
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Response
// @Failure 404 {object} model.Response
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := service.DeleteUser(id); err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}


