package incoming

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/dtos/validators"
	"github.com/go-playground/validator"
)

type YoutubeVideoDownloadDTO struct {
	VideoUrl string `json:"VideoUrl" validate:"required,url"`

	// Accept either itag or quality
	Itag    int    `json:"itag" validate:"omitempty,gt=0"`
	Quality string `json:"quality" validate:"omitempty,quality"`
}

// ValidateData validates the fields in the DTO
func (data *YoutubeVideoDownloadDTO) ValidateData() error {
	validate := validator.New()

	// Register custom validation for "quality"
	validate.RegisterValidation("quality", validators.QualityValidator)

	// Validate the DTO
	err := validate.Struct(data)
	if err != nil {
		return err // Return validation errors
	}
	return nil // Return nil if validation succeeded
}
