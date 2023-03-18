package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type GenerateTemplate struct {
	TemplateID string     `json:"template_id,omitempty"`
	Title      string     `json:"title,omitempty"`
	Variables  []Variable `json:"variables,omitempty"`
	Test       bool       `json:"test,omitempty"`
}

type Variable struct {
	Name       string       `json:"name,omitempty"`
	Properties []Properties `json:"properties,omitempty"`
}

type Properties struct {
	Default string `json:"default,omitempty"`
	Name    string `json:"name,omitempty"`
}

func (t *GenerateTemplate) Path() (string, error) {
	return "/template.generate", nil
}

func (t *GenerateTemplate) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (t *GenerateTemplate) ContentType() string {
	return "application/json"
}
