package models

import "time"

type Task struct {
	ID          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	DueDate     time.Time `bson:"due_date"`
	Status      string    `bson:"status"`
}

