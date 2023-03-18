package response

import (
	"io"
)

type ListAvatars struct {
	Response
	Data struct {
		Avatars []Avatar `json:"avatars"`
	} `json:"data"`
}

type Avatar struct {
	AvatarID     int           `json:"avatar_id"`
	AvatarStates []AvatarState `json:"avatar_states"`
	CreatedAt    int           `json:"created_at"`
	Gender       string        `json:"gender"`
	Name         string        `json:"name"`
}

type AvatarState struct {
	AvatarRace           string                   `json:"avatar_race"`
	AvatarStyle          string                   `json:"avatar_style"`
	CirclePreview        string                   `json:"circle_preview"`
	CirclePreviewHeight  int                      `json:"circle_preview_height"`
	CirclePreviewWidth   int                      `json:"circle_preview_width"`
	CircleThumbnailSmall string                   `json:"circle_thumbnail_small"`
	CloseUpCropX         int                      `json:"close_up_crop_x"`
	CloseUpPreview       string                   `json:"close_up_preview"`
	CloseUpPreviewHeight int                      `json:"close_up_preview_height"`
	CloseUpPreviewWidth  int                      `json:"close_up_preview_width"`
	CreatedAt            int                      `json:"created_at"`
	CropX                int                      `json:"crop_x,omitempty"`
	CustomAvatarType     string                   `json:"custom_avatar_type"`
	EndSpeechOffset      float64                  `json:"end_speech_offset"`
	FavoriteUpdatedAt    int                      `json:"favorite_updated_at"`
	Gender               string                   `json:"gender"`
	ID                   string                   `json:"id"`
	ImageURL             string                   `json:"image_url,omitempty"`
	IsCustomer           bool                     `json:"is_customer"`
	IsFavorite           bool                     `json:"is_favorite"`
	IsPreset             bool                     `json:"is_preset"`
	Name                 string                   `json:"name"`
	NormalPreviewHeight  int                      `json:"normal_preview_height,omitempty"`
	NormalPreviewWidth   int                      `json:"normal_preview_width,omitempty"`
	PoseName             string                   `json:"pose_name"`
	SpecialVoice         SpecialVoice             `json:"special_voice"`
	StartSpeechOffset    float64                  `json:"start_speech_offset"`
	SupportMatting       bool                     `json:"support_matting"`
	SupportedTemplates   []map[string]interface{} `json:"supported_templates"`
	ThumbnailAlpha       string                   `json:"thumbnail_alpha"`
	ThumbnailSmall       string                   `json:"thumbnail_small"`
	VideoURL             VideoURL                 `json:"video_url"`
}

type SpecialVoice struct {
	Pitch         int     `json:"pitch"`
	Speed         float64 `json:"speed"`
	Text          string  `json:"text"`
	TtsAudioURL   string  `json:"tts_audio_url"`
	TtsDuration   float64 `json:"tts_duration"`
	VoiceGender   string  `json:"voice_gender"`
	VoiceID       string  `json:"voice_id"`
	VoiceLanguage string  `json:"voice_language"`
}

type VideoURL struct {
	Grey string `json:"grey"`
}

func (a *ListAvatars) Decode(body io.Reader) error {
	return decode(body, a)
}

func (a *ListAvatars) AcceptContentType() string {
	return "application/json"
}
