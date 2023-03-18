package response

import (
	"io"
)

type GenerateVideo struct {
	Response
	Data struct {
		VideoID string `json:"video_id"`
	} `json:"data"`
}

func (v *GenerateVideo) Decode(body io.Reader) error {
	return decode(body, v)
}

func (v *GenerateVideo) AcceptContentType() string {
	return "application/json"
}
