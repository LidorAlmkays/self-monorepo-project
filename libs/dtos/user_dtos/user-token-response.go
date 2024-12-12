package user_dtos

type UserTokenResponseDTO struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
