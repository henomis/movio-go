package response

import (
	"io"
)

type ListVoices struct {
	Response
	Data struct {
		Voices []Voice `json:"voices"`
	} `json:"data"`
}

type Voice struct {
	VoiceID  string `json:"voice_id"`
	Gender   string `json:"gender"`
	Language string `json:"language"`
	Locale   string `json:"locale"`
	Name     string `json:"name"`
	Preview  string `json:"preview"`
}

func (v *ListVoices) Decode(body io.Reader) error {
	return decode(body, v)
}

func (v *ListVoices) AcceptContentType() string {
	return "application/json"
}
