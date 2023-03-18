package response

import (
	"io"
)

type GenerateTemplate struct {
	Response
	Data struct {
		VideoID string `json:"video_id"`
	} `json:"data"`
}

func (c *GenerateTemplate) Decode(body io.Reader) error {
	return decode(body, c)
}

func (a *GenerateTemplate) AcceptContentType() string {
	return "application/json"
}
