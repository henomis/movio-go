package request

import (
	"io"

	"github.com/henomis/restclientgo"
)

type GetTemplate struct {
	VideoID string `json:"video_id"`
}

func (t *GetTemplate) Path() (string, error) {
	return "/template.get?" + t.getParameters().Encode(), nil
}

func (t *GetTemplate) getParameters() *restclientgo.URLValues {
	parameters := restclientgo.NewURLValues()
	parameters.Add("video_id", &t.VideoID)
	return parameters
}

func (t *GetTemplate) Encode() (io.Reader, error) {
	return nil, nil
}

func (t *GetTemplate) ContentType() string {
	return ""
}
