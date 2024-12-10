package user

import (
	"time"

	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/enums"
)

type AddUserDTO struct {
	Name     string     `json:"name"`
	UserName string     `json:"username"`
	Password string     `json:"password"`
	Email    string     `json:"email"`
	BirthDay time.Time  `json:"birthDay"`
	Role     enums.Role `json:"role"`
}
