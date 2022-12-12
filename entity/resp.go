package entity

type Response struct {
	Message        ResponseMessage `json:"message"`
	ConversationID string          `json:"conversation_id"`
	Error          any             `json:"error"`
}

type ResponseMessage struct {
	ID         string           `json:"id"`
	Role       string           `json:"role"`
	User       any              `json:"user"`
	CreateTime any              `json:"create_time"`
	UpdateTime any              `json:"update_time"`
	Content    ResponseContent  `json:"content"`
	EndTurn    any              `json:"end_turn"`
	Weight     float64          `json:"weight"`
	Metadata   ResponseMetadata `json:"metadata"`
	Recipient  string           `json:"recipient"`
}

type ResponseContent struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

type ResponseMetadata struct {
}
