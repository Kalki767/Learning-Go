package initializers

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
func CreateConnection() {
	// Connect to the database
	dbUri := os.Getenv("DB_URI")
	clientOption := options.Client().ApplyURI(dbUri)

	
	client, err := mongo.Connect(context.TODO(), clientOption)

	// Check the connection
	if err != nil {
		log.Fatal("MongoDB Connection Failed:", err)
		return 
	}
	
	// Ping the database to check if the connection is successful
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("MongoDB Ping Failed:", err)
		return 
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
	fmt.Println("Client:", Client)
	
}

