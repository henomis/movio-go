package request

import "io"

type ListTalkingPhotos struct{}

func (t *ListTalkingPhotos) Path() (string, error) {
	return "/talking_photo.list", nil
}

func (t *ListTalkingPhotos) Encode() (io.Reader, error) {
	return nil, nil
}

func (t *ListTalkingPhotos) ContentType() string {
	return ""
}
