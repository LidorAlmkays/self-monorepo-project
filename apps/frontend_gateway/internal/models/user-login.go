package models

type UserLoginModel struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
