package validators

import "github.com/go-playground/validator"

// Custom Validation Function
func QualityValidator(fl validator.FieldLevel) bool {
	validQualities := []int{144, 240, 360, 480, 720, 1080, 1440, 2160}
	value := fl.Field().Int()
	for _, q := range validQualities {
		if value == int64(q) {
			return true
		}
	}
	return false
}
