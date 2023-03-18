package response

import (
	"encoding/json"
	"fmt"
	"io"
)

type Response struct {
	StatusCode int    `json:"-"`
	Code       int    `json:"code"`
	Message    string `json:"message"`
	RawBody    string `json:"-"`
}

func (s *Response) IsSuccess() bool {
	return (s.StatusCode >= 200 && s.StatusCode < 300)
}

func (r *Response) Error() error {
	if len(r.RawBody) > 0 {
		return fmt.Errorf("[%d]: %s", r.StatusCode, r.RawBody)
	}
	return fmt.Errorf("[%d]: %s", r.StatusCode, r.Message)
}

func (r *Response) SetStatusCode(code int) error {
	r.StatusCode = code
	return nil
}

func (r *Response) SetBody(body io.Reader) error {
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	r.RawBody = string(bodyBytes)
	return nil
}

func decode(body io.Reader, data interface{}) error {
	return json.NewDecoder(body).Decode(data)
}
