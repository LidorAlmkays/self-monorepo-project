package models

import "time"

type UserModel struct {
	Name     string    `json:"name" bson:"name" validate:"required"`
	UserName string    `json:"username" bson:"username" validate:"required"`
	Password string    `json:"password" bson:"password" validate:"required"`
	BirthDay time.Time `json:"birthDay" bson:"Birth_day" validate:"required"`
	Email    string    `json:"email" bson:"email" validate:"required"`
}
