package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/handler"
)

func LoadRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Auth routes
		api.POST("/users/register", handler.RegisterUser)
		api.POST("/users/login", handler.LoginUser)
		api.POST("/users/reset-password", handler.RequestPasswordReset)
		api.POST("/users/reset-password-confirm", handler.ConfirmPasswordReset)

		// User CRUD routes
		api.GET("/users/:id", handler.GetUser)
		api.PATCH("/users/:id", handler.UpdateUser)
		api.DELETE("/users/:id", handler.DeleteUser)
	}
}
