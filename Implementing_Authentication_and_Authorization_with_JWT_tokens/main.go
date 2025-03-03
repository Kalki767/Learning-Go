package main

import (
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/controllers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/data"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/initializers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/routers"
)


func init() {

	//Load environment variables
	initializers.LoadEnvVariables()
	initializers.CreateConnection()
	controllers.InitializeUserDB()
	data.InitializeTaskDB()
}
func main() {

	
	routers.SetRouter()

}