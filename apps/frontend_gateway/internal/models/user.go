package models

import "time"

type UserModel struct {
	Name     string    `json:"Name"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	BirthDay time.Time `json:"age"`
}
