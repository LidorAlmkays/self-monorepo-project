package user

type AuthenticateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
