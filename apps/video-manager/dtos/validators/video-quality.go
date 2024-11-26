package validators

import "github.com/go-playground/validator"

// QualityValidator ensures the quality is valid as either a YouTube quality string or a number
func QualityValidator(fl validator.FieldLevel) bool {
	// All possible YouTube quality types
	var allowedQualities = map[string]struct{}{
		"144p":   {},
		"240p":   {},
		"360p":   {},
		"480p":   {},
		"720p":   {},
		"1080p":  {},
		"1440p":  {},
		"2160p":  {},
		"hd1080": {},
		"hd720":  {},
		"hd1440": {},
		"hd2160": {},
		"large":  {},
		"medium": {},
		"small":  {},
	}

	// Check if the quality is in the allowed qualities
	quality := fl.Field().String()
	_, exists := allowedQualities[quality]
	return exists

}
