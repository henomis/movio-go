package request

import (
	"bytes"
	"encoding/json"
	"io"
)

type AvatarStyle string

const (
	AvatarStyleCircle AvatarStyle = "circle"
	AvatarStyleNormal AvatarStyle = "normal"
)

type GenerateVideo struct {
	Background string `json:"background,omitempty"`
	Ratio      string `json:"ratio,omitempty"`
	Test       bool   `json:"test,omitempty"`
	Version    string `json:"version,omitempty"`
	Clips      []Clip `json:"clips,omitempty"`
}

type Clip struct {
	AvatarID          string      `json:"avatar_id,omitempty"`
	VoiceID           string      `json:"voice_id,omitempty"`
	AvatarStyle       AvatarStyle `json:"avatar_style,omitempty"`
	InputText         string      `json:"input_text,omitempty"`
	InputAudio        string      `json:"-"` // This field is not documented in the API docs
	Scale             int         `json:"scale,omitempty"`
	Offset            *Offset     `json:"offset,omitempty"`
	TalkingPhotoID    string      `json:"talking_photo_id,omitempty"`
	TalkingPhotoStyle AvatarStyle `json:"talking_photo_style,omitempty"`
}

type Offset struct {
	X int `json:"x,omitempty"`
	Y int `json:"y,omitempty"`
}

func (v *GenerateVideo) Path() (string, error) {
	return "/video.generate", nil
}

func (v *GenerateVideo) Encode() (io.Reader, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(jsonBytes), nil
}

func (v *GenerateVideo) ContentType() string {
	return "application/json"
}
