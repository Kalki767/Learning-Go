package data

import (
	"context"
	"log"
	

	"github.com/Kalki767/Task_Manager_API_With_MongoDB/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func CreateConnection() *mongo.Client {
	// Connect to the database
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017/")
	client, err := mongo.Connect(context.TODO(), clientOption)

	// Check the connection
	if err != nil {
		log.Fatal(err)
		return nil
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return client
}

var client = CreateConnection()

var taskCollection = client.Database("Task_database").Collection("tasks")
var userCollection = client.Database("User_database").Collection("users")

// Get all tasks
func GetAllTasks()[]models.Task {
	// hold a cursor to iterate through the tasks
	cursor, err := taskCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	// Close the cursor once the function is done
	defer cursor.Close(context.TODO())

	// Create a slice to hold the tasks
	var Tasks []models.Task

	// Iterate through the cursor and decode the tasks
	for cursor.Next(context.TODO()) {
		var task models.Task
		cursor.Decode(&task)
		Tasks = append(Tasks, task)
	}
	return Tasks
}

// Get a task by ID
func GetTask(id string) models.Task{

	var task models.Task

	// Find the task with the given ID
	err := taskCollection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&task)
	if err != nil {
		log.Fatal(err)
		return models.Task{}
	}
	// Return the task if the task is  found
	return task
}

// Create a new task
func PostTask(newTask models.Task) error {
	
	// Insert the new task
	_, err := taskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	return nil
}

// Update a task by ID
func UpdateTask(id string, updatedTask models.Task) error{

	// Create an update filter
	update := bson.M{
		"$set": bson.M{
			"title": updatedTask.Title,
			"description": updatedTask.Description,
			"due_date": updatedTask.DueDate,
			"status": updatedTask.Status,
		},
	}

	// Update the task with the given ID
	_, err := taskCollection.UpdateOne(context.TODO(), bson.M{"_id": id},update)
	if err != nil {
		log.Fatal(err)
		return err
	}
	
	return nil
}

// Delete a task by ID
func DeleteTask(id string) error{

	// Delete the task with the given ID
	_, err := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

// func RegisterUser(user models.User) error{

// }
