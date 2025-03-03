package data

import (
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/initializers"
	"Implementing_Authentication_and_Authorization_with_JWT_tokens/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var taskCollection *mongo.Collection
func InitializeTaskDB() {

	if initializers.Client == nil {
		panic("MongoDB Client is not initialized! Ensure ConnectMongoDB() is called.")
	}
	taskCollection = initializers.Client.Database("Task_database").Collection("tasks")
}


// Get all tasks
func GetAllTasks() [] models.Task {
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
func GetTask(id string) models.Task {

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
func UpdateTask(id string, updatedTask models.Task) error {

	// Create an update filter
	update := bson.M{
		"$set": bson.M{
			"title":       updatedTask.Title,
			"description": updatedTask.Description,
			"due_date":    updatedTask.DueDate,
			"status":      updatedTask.Status,
		},
	}

	// Update the task with the given ID
	_, err := taskCollection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

// Delete a task by ID
func DeleteTask(id string) error {

	// Delete the task with the given ID
	_, err := taskCollection.DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
