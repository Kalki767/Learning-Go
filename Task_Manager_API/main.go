package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate time.Time `json:"due_date"`
	Status string `json:"status"`

}

// Mock data for tasks
var tasks = []Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// Get all tasks
func getAllTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, tasks)
}

// Get a task by ID
func getTask(context *gin.Context) {
	id := context.Param("id")

	// Loop through tasks to find the task with the given ID
	for _, task := range tasks{
		if task.ID == id{
			context.IndentedJSON(http.StatusOK, task)
			return
		}
	}

	// Return an error message if the task is not found
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

// Create a new task
func postTask(context *gin.Context) {
	var newTask Task

	// Bind the JSON data from the request body to the newTask variable
	if err := context.BindJSON(&newTask); err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Append the task to the tasks slice
	tasks = append(tasks, newTask)
	context.IndentedJSON(http.StatusCreated, newTask)
}

// Update a task by ID
func updateTask(context *gin.Context) {
	id := context.Param("id")

	var updatedTask Task

	// Bind the JSON data from the request body to the updatedTask variable
	if err := context.ShouldBindJSON(&updatedTask); err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Loop through tasks to find the task with the given ID
	for i, task := range tasks{
		if task.ID == id{
			if updatedTask.Title != ""{
				tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				tasks[i].Description = updatedTask.Description
			}
			if task.Status != "" {
				task.Status = updatedTask.Status
			}
			context.IndentedJSON(http.StatusOK, gin.H{"message" : "Task Updated"})
			return
		}
	}

	// Return an error message if the task is not found
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Delete a task by ID
func deleteTask(context *gin.Context) {
	id := context.Param("id")

	// Loop through tasks to find the task with the given ID
	for i, task := range tasks{
		if task.ID == id{
			tasks = append(tasks[:i],tasks[i+1:]...)
			context.IndentedJSON(http.StatusOK, gin.H{"message" : "Task deleted Successfully"})
			return
		}
	}

	// Return an error message if the task is not found
	context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task not found"})
}

// Main function to run the application
func main() {
	router := gin.Default()

	router.GET("/tasks", getAllTasks)
	router.GET("/tasks/:id", getTask)
	router.POST("/tasks", postTask)
	router.PUT("/tasks/:id", updateTask)
	router.DELETE("/tasks/:id", deleteTask)

	router.Run()
}