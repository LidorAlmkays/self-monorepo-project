package user_dtos

import "time"

type RequestToAddUserDTO struct {
	Name     string    `json:"name"`
	UserName string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	BirthDay time.Time `json:"birthDay"`
}
