package data

import (
	"errors"
	"time"
	"github.com/Kalki767/Task_Manager_API/models"
)

// Mock data for tasks
var Tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// Get all tasks
func GetAllTasks()[]models.Task {
	return Tasks
}

// Get a task by ID
func GetTask(id string) models.Task{

	// Loop through tasks to find the task with the given ID
	for _, task := range Tasks{
		if task.ID == id{
			return task
		}
	}

	// Return an empty task if the task is not found
	return models.Task{}
}

// Create a new task
func PostTask(newTask models.Task)  {
	
	// Append the task to the tasks slice
	Tasks = append(Tasks, newTask)
	
}

// Update a task by ID
func UpdateTask(id string, updatedTask models.Task) error{

	// Loop through tasks to find the task with the given ID
	for i, task := range Tasks{
		if task.ID == id{
			if updatedTask.Title != ""{
				Tasks[i].Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				Tasks[i].Description = updatedTask.Description
			}
			if task.Status != "" {
				Tasks[i].Status = updatedTask.Status
			}
			
			return nil
		}
	}

	// Return an error message if the task is not found
	return errors.New("task not found")
}

// Delete a task by ID
func DeleteTask(id string) error{

	// Loop through tasks to find the task with the given ID
	for i, task := range Tasks{
		if task.ID == id{
			Tasks = append(Tasks[:i],Tasks[i+1:]...)
			return nil
		}
	}

	// Return an error message if the task is not found
	return errors.New("task not found")
}
