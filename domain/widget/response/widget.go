package response

type WidgetSetting struct {
	ID              string `json:"id"`
	UserID          string `json:"user_id"`
	TTSEnabled      bool   `json:"tts_enabled"`
	NSFWFilter      bool   `json:"nsfw_filter"`
	MessageDuration int    `json:"message_duration"`
	MinAmount       int64  `json:"min_amount"`
}
