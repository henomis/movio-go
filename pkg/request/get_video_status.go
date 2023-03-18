package request

import (
	"io"

	"github.com/henomis/restclientgo"
)

type GetVideoStatus struct {
	VideoID string `json:"video_id"`
}

func (v *GetVideoStatus) Path() (string, error) {
	return "/video_status.get?" + v.getParameters().Encode(), nil
}

func (v *GetVideoStatus) getParameters() *restclientgo.URLValues {
	parameters := restclientgo.NewURLValues()
	parameters.Add("video_id", &v.VideoID)
	return parameters
}

func (v *GetVideoStatus) Encode() (io.Reader, error) {
	return nil, nil
}

func (v *GetVideoStatus) ContentType() string {
	return ""
}
