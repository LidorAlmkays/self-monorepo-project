package incoming

import (
	"github.com/LidorAlmkays/self-monorepo-project/apps/video-manager/dtos/validators"
	"github.com/go-playground/validator"
)

type YoutubeVideoDownloadDTO struct {
	VideoUrl     string `json:"VideoUrl" validate:"required,url"`
	VideoQuality int    `json:"VideoQuality" validate:"quality"`
}

func (data *YoutubeVideoDownloadDTO) ValidateData() error {
	// Validate the Struct
	validate := validator.New()
	// Register Custom Validation
	validate.RegisterValidation("quality", validators.QualityValidator)
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
