package user

type UserTokenResponseDTO struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
