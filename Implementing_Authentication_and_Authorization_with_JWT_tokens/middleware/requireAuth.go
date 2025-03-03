package middleware

import (
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/controllers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RequireAuth(c *gin.Context) {

	//Get the cookie off the req
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		fmt.Println("Error getting token: ", err)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate the token
	// Parse takes the token string and a function for looking up the key. The latter is especially
	
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		fmt.Println("Error parsing token: ")
		log.Fatal(err)
	}

	
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		
		//check the expiration
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			fmt.Println("Token expired")
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Convert `claims["id"]` from string to `primitive.ObjectID`
		userIDStr, ok := claims["id"].(string)
		if !ok {
			fmt.Println("Invalid user ID format in token")
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Convert userIDStr to ObjectID
		userID, err := primitive.ObjectIDFromHex(userIDStr)
		if err != nil {
			fmt.Println("Invalid ObjectID: ", userIDStr)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		//Find the use with the user id
		var user models.User
		if err := controllers.UserDB.FindOne(context.TODO(), bson.M{"_id": userID}).Decode(&user); err != nil {
			fmt.Println("Error finding user: ", claims["id"])
			c.AbortWithStatus(http.StatusUnauthorized)
		}


		//Attach to request
		c.Set("user", user)

		//Continue

		c.Next()

	} else {
		fmt.Println("final else: ")
		c.AbortWithStatus(http.StatusUnauthorized)
	}


	
}