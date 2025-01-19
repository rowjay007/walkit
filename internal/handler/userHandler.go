package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/rowjay007/walkit/internal/model"
    "github.com/rowjay007/walkit/internal/service"
    "github.com/rowjay007/walkit/pkg/util"
)

func GetMe(c *gin.Context) {
    userID := c.GetString("userID")
    user, err := service.GetUser(userID)
    if err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, user)
}

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

// DeleteMe deletes the current user's account
func DeleteMe(c *gin.Context) {
    userID := c.GetString("userID")
    if err := service.DeleteUser(userID); err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "Account deleted successfully"})
}

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

func GetUser(c *gin.Context) {
    id := c.Param("id")
    user, err := service.GetUser(id)
    if err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, user)
}

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

// DeleteUser deletes a specific user
func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := service.DeleteUser(id); err != nil {
        util.RespondWithError(c, http.StatusInternalServerError, err.Error())
        return
    }
    util.RespondWithJSON(c, http.StatusOK, gin.H{"message": "User deleted successfully"})
}


