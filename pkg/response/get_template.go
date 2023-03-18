package response

import (
	"io"
)

type GetTemplate struct {
	Response
	Data struct {
		Scenes     []Scenes `json:"scenes"`
		TemplateID string   `json:"template_id"`
		VideoID    string   `json:"video_id"`
	} `json:"data"`
}

type Scenes struct {
	Variables []Variable `json:"variables"`
}

type Variable struct {
	Name       string       `json:"name"`
	Properties []Properties `json:"properties"`
	Type       string       `json:"type"`
}

type Properties struct {
	Default string `json:"default"`
	Name    string `json:"name"`
}

func (t *GetTemplate) Decode(body io.Reader) error {
	return decode(body, t)
}

func (t *GetTemplate) AcceptContentType() string {
	return "application/json"
}
