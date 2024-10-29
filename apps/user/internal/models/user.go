package models

import "time"

type UserModel struct {
	Name     string    `bson:"Name"`
	UserName string    `bson:"user_name"`
	Password string    `bson:"password"`
	BirthDay time.Time `bson:"BirthDay"`
	Email    string    `bson:"email"`
}
