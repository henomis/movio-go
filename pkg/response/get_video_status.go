package response

import (
	"io"
)

type GetVideoStatus struct {
	Response
	Data struct {
		ID       string `json:"id"`
		Error    string `json:"error"`
		Status   string `json:"status"`
		VideoURL string `json:"video_url"`
	} `json:"data"`
}

func (v *GetVideoStatus) Decode(body io.Reader) error {
	return decode(body, v)
}

func (v *GetVideoStatus) AcceptContentType() string {
	return "application/json"
}
