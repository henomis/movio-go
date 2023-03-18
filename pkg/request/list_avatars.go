package request

import "io"

type ListAvatars struct{}

func (a *ListAvatars) Path() (string, error) {
	return "/avatar.list", nil
}

func (a *ListAvatars) Encode() (io.Reader, error) {
	return nil, nil
}

func (a *ListAvatars) ContentType() string {
	return ""
}
