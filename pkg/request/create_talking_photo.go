package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type CreateTalkingPhoto struct {
	StorageKey string `json:"storage_key"`
}

func (t *CreateTalkingPhoto) Path() (string, error) {
	return "/talking_photo.create", nil
}

func (t *CreateTalkingPhoto) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (t *CreateTalkingPhoto) ContentType() string {
	return "application/json"
}
