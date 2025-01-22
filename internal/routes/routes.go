package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rowjay007/walkit/internal/handler"
	"github.com/rowjay007/walkit/internal/middleware"
)

func LoadRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", handler.RegisterUser)
			auth.POST("/login", handler.LoginUser)
			auth.POST("/forgot-password", handler.RequestPasswordReset)
			auth.POST("/reset-password", handler.ConfirmPasswordReset)
		}

		users := api.Group("/users")
		users.Use(middleware.JWTAuthMiddleware)
		{
			users.GET("/me", handler.GetMe)
			users.PATCH("/me", handler.UpdateMe)
			users.DELETE("/me", handler.DeleteMe)

			users.GET("", handler.ListUsers)
			users.GET("/:id", handler.GetUser)
			users.PATCH("/:id", handler.UpdateUser)
			users.DELETE("/:id", handler.DeleteUser)
		}

		exercises := api.Group("/exercises")
		exercises.Use(middleware.JWTAuthMiddleware)

		{
			exercises.POST("", handler.CreateExercise)
			exercises.GET("", handler.ListExercises)
			exercises.GET("/:id", handler.GetExercise)
			exercises.PATCH("/:id", handler.UpdateExercise)
			exercises.DELETE("/:id", handler.DeleteExercise)
		}

		workouts := api.Group("/workouts")
		workouts.Use(middleware.JWTAuthMiddleware)

		{
			workouts.POST("", handler.CreateWorkout)
			workouts.GET("/:id", handler.GetWorkout)
			workouts.PATCH("/:id", handler.UpdateWorkout)
			workouts.DELETE("/:id", handler.DeleteWorkout)
		}

	}
}
