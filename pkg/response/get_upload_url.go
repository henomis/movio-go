package response

import (
	"io"
)

type GetUploadURL struct {
	Response
	Data struct {
		StorageKey string `json:"storage_key"`
		UploadURL  string `json:"upload_url"`
	} `json:"data"`
}

func (u *GetUploadURL) Decode(body io.Reader) error {
	return decode(body, u)
}

func (u *GetUploadURL) AcceptContentType() string {
	return "application/json"
}
