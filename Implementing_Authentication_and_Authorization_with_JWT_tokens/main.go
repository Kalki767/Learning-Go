package main

import (
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/controllers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/initializers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/middleware"

	"github.com/gin-gonic/gin"
)


func init() {

	//Load environment variables
	initializers.LoadEnvVariables()
	initializers.CreateConnection()
	controllers.InitializeUserDB()
}
func main() {

	
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)


	router.Run()

}