package controllers

import (
	"net/http"

	"github.com/Kalki767/Task_Manager_API/data"
	"github.com/Kalki767/Task_Manager_API/models"
	"github.com/gin-gonic/gin"
)

// Get all tasks
func GetAllTasks(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.GetAllTasks())
}

// Get a task by ID
func GetTask(context *gin.Context) {
	id := context.Param("id")

	task := data.GetTask(id)
	if task.ID != "" {
		context.IndentedJSON(http.StatusOK, task)
		return
	}

	// Return an error message if the task is not found
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}


// Create a new task
func PostTask(context *gin.Context) {
	var newTask models.Task

	// Bind the JSON data from the request body to the newTask variable
	if err := context.BindJSON(&newTask); err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Append the task to the tasks slice
	data.PostTask(newTask)
	context.IndentedJSON(http.StatusCreated, newTask)
}

// Update a task by ID
func UpdateTask(context *gin.Context) {
	id := context.Param("id")

	var updatedTask models.Task

	// Bind the JSON data from the request body to the updatedTask variable
	if err := context.ShouldBindJSON(&updatedTask); err != nil{
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := data.UpdateTask(id, updatedTask)
	if err == nil {
		context.IndentedJSON(http.StatusOK, gin.H{"message" : "Task Updated"})
		return
	}

	// Return an error message if the task is not found
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// Delete a task by ID
func DeleteTask(context *gin.Context) {
	id := context.Param("id")

	err := data.DeleteTask(id)
	if err == nil {
		context.IndentedJSON(http.StatusOK, gin.H{"message" : "Task deleted Successfully"})
		return
	}
	
	// Return an error message if the task is not found
	context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Task not found"})
}
