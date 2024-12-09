package validators

import (
	"github.com/LidorAlmkays/self-monorepo-project/libs/golang/enums"
	"github.com/go-playground/validator"
)

func ProgramModeValidator(fl validator.FieldLevel) bool {
	var value enums.ProgramMode = enums.ProgramMode(fl.Field().String())
	return value.IsValid()
}
