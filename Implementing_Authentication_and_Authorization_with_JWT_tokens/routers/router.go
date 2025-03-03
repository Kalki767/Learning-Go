package routers

import (
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/controllers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/middleware"

	"github.com/gin-gonic/gin"
)

func SetRouter() {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks:id", controllers.GetTask)
	router.POST("/tasks", controllers.PostTask)
	router.PUT("/tasks:id", controllers.UpdateTask)
	router.DELETE("/tasks:id", controllers.DeleteTask)

	router.Run()
}