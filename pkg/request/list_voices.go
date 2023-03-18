package request

import "io"

type ListVoices struct{}

func (v *ListVoices) Path() (string, error) {
	return "/voice.list", nil
}

func (v *ListVoices) Encode() (io.Reader, error) {
	return nil, nil
}

func (v *ListVoices) ContentType() string {
	return ""
}
