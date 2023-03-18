package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type GetUploadURL struct {
	URLContentType string `json:"content_type"`
}

func (u *GetUploadURL) Path() (string, error) {
	return "/upload_url.get", nil
}

func (u *GetUploadURL) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (u *GetUploadURL) ContentType() string {
	return "application/json"
}
