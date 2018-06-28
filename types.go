package tg

type Update struct {
	UpdateID       int           `json:"update_id"`
	Message        Message       `json:"message"`
	ReplyToMessage Message       `json:"reply_to_message"`
	CallBackQuery  CallBackQuery `json:"callback_query"`
}

type CallBackQuery struct {
	ID           string  `json:"id"`
	Message      Message `json:"message"`
	From         User    `json:"from"`
	ChatInstance string  `json:"chat_instance"`
	Data         string  `json:"data"`
}

type Message struct {
	MessageID int    `json:"message_id"`
	Date      int    `json:"date"`
	Text      string `json:"text" validate:"required"`
	Caption   string `json:"caption"`
	From      User   `json:"from" validate:"required"`
	Chat      Chat   `json:"chat" validate:"required"`
}

type User struct {
	ID           int    `json:"id" validate:"required"`
	IsBot        bool   `json:"is_bit"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username" validate:"required"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID          int    `json:"id" validate:"required"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type SendMessageRequest struct {
	ChatID int
	Text   string
}

const (
	ActionTyping          = "typing"
	ActionUploadPhoto     = "upload_photo"
	ActionRecordVideo     = "record_video"
	ActionUploadVideo     = "upload_video"
	ActionRecordAudio     = "record_audio"
	ActionUploadAudio     = "upload_audio"
	ActionUploadDocument  = "upload_document"
	ActionFindLocation    = "find_location"
	ActionRecordVideoNote = "record_video_note"
	ActionUploadVideoNote = "upload_video_note"
)

type ImageInlineRequest struct {
	ChatID      int                  `json:"chat_id"`
	Photo       string               `json:"photo"`
	ReplyMarkup InlineKeyboardMarkup `json:"reply_markup"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}
