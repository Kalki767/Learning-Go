package router

import (
	"github.com/gin-gonic/gin"
	"github.com/Kalki767/Task_Manager_API_With_MongoDB/controllers"
)

func SetRouter() {
	router := gin.Default()

	router.GET("/tasks", controllers.GetAllTasks)
	router.GET("/tasks/:id", controllers.GetTask)
	router.POST("/tasks", controllers.PostTask)
	router.PUT("/tasks/:id", controllers.UpdateTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)

	router.Run()
}