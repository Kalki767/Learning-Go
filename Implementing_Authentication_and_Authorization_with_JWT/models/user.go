package models

type User struct {
	ID uint `bson:"_id"`
	Email string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"-"`
}