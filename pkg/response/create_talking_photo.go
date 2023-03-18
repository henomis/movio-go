package response

import (
	"io"
)

type CreateTalkingPhoto struct {
	Response
	Data struct {
		ID                 string `json:"id"`
		ImageHeight        int    `json:"image_height"`
		ImageWidth         int    `json:"image_width"`
		ImageURL           string `json:"image_url"`
		EnhancedImageURL   string `json:"enhanced_image_url"`
		EnhancedCroppedURL string `json:"enhanced_cropped_url"`
		CroppedURL         string `json:"cropped_url"`
		CroppedHeight      int    `json:"cropped_height"`
		CroppedWidth       int    `json:"cropped_width"`
	} `json:"data"`
}

func (t *CreateTalkingPhoto) Decode(body io.Reader) error {
	return decode(body, t)
}

func (t *CreateTalkingPhoto) AcceptContentType() string {
	return "application/json"
}
