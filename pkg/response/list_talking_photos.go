package response

import (
	"io"
)

type ListTalkingPhotos struct {
	Response
	Data []struct {
		ID          string `json:"id"`
		CircleImage string `json:"circle_image"`
		ImageURL    string `json:"image_url"`
	} `json:"data"`
}

func (t *ListTalkingPhotos) Decode(body io.Reader) error {
	return decode(body, t)
}

func (t *ListTalkingPhotos) AcceptContentType() string {
	return "application/json"
}
