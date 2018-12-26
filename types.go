package tg

type Update struct {
	UpdateID       int64         `json:"update_id"`
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
	MessageID int64  `json:"message_id"`
	Date      int64  `json:"date"`
	Text      string `json:"text" validate:"required"`
	Caption   string `json:"caption"`
	From      User   `json:"from" validate:"required"`
	Chat      Chat   `json:"chat" validate:"required"`
}

type User struct {
	ID           int64  `json:"id" validate:"required"`
	IsBot        bool   `json:"is_bit"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username" validate:"required"`
	LanguageCode string `json:"language_code"`
}

type Chat struct {
	ID          int64  `json:"id" validate:"required"`
	Type        string `json:"type"`
	Title       string `json:"title"`
	Username    string `json:"username"`
	Description string `json:"description"`
}

type SendMessageRequest struct {
	ChatID int64
	Text   string
}

type SendMessageMarkupRequest struct {
	ChatID      int64       `json:"chat_id"`
	Text        string      `json:"text"`
	ReplyMarkup interface{} `json:"reply_markup"`
	ParseMode   string      `json:"parse_mode"`
}

type Photo struct {
	FileID   string `json:"file_id"`
	FileSize int64  `json:"file_size"`
	Width    int64  `json:"width"`
	Height   int64  `json:"height"`
}

type PhotoResponse struct {
	OK     bool        `json:"ok"`
	Result PhotoResult `json:"result"`
}

type PhotoResult struct {
	Photo []Photo `json:"photo"`
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
	ChatID              int64                `json:"chat_id"`
	Photo               string               `json:"photo"`
	DisableNotification bool                 `json:"disable_notification"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
	Caption             string               `json:"caption"`
	ParseMode           string               `json:"parse_mode"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

type ReplyKeyboardMarkup struct {
	Keyboard        [][]KeyboardButton `json:"keyboard"`
	ResizeKeyboard  bool               `json:"resize_keyboard"`
	OneTimeKeyboard bool               `json:"one_time_keyboard"`
	Selective       bool               `json:"selective"`
}

type InlineKeyboardButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}

type KeyboardButton struct {
	Text            string `json:"text"`
	RequestContact  bool   `json:"request_contact"`
	RequestLocation bool   `json:"request_location"`
}

type EditMessageCaptionRequest struct {
	ChatID    int64  `json:"chat_id"`
	MessageID int64  `json:"message_id"`
	Caption   string `json:"caption"`
	ParseMode string `json:"parse_mode"`
}
