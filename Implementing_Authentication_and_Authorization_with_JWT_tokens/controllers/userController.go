package controllers

import (
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/initializers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/models"
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var UserDB *mongo.Collection
func InitializeUserDB() {

	if initializers.Client == nil {
		panic("MongoDB Client is not initialized! Ensure ConnectMongoDB() is called.")
	}
	UserDB = initializers.Client.Database("User_database").Collection("users")
}


func SignUp(c *gin.Context){

	//Get email and password from the request
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
		return
	}

	//Check if the email is already registered
	var result bson.M
	err := UserDB.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&result)

	if err == nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is already registered"})
		return
	}


	//Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate password"})
		return
	}


	//Create a new user
	user.Password = string(hashedPassword)
	_, err = UserDB.InsertOne(context.TODO(), user)
	if err != nil {
		fmt.Println("Error inserting user: ", err)
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	//Respond with a success message
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {

	//Get email and password from the request
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})

		return
	}


	//look for the user in the database
	var databaseUser models.User
	err := UserDB.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&databaseUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}

	//compare the password with the hashed password
	err = bcrypt.CompareHashAndPassword([]byte(databaseUser.Password), []byte(user.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email or password"})
		return
	}


	//generate a JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": databaseUser.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	secret_key := os.Getenv("SECRET_KEY")
	tokenString, err := token.SignedString([]byte(secret_key))

	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}

	//store it in cookie
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600 * 24 * 30, "", "", false, true)

}

func Validate(c *gin.Context) {
	c.JSON(200, gin.H{"message": "User is authenticated"})
}