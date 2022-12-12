package entity

import (
	"github.com/golang-infrastructure/go-pointer"
	"github.com/google/uuid"
)

type Request struct {
	Action          string           `json:"action"`
	Messages        []RequestMessage `json:"messages"`
	ConversationID  *string          `json:"conversation_id"`
	ParentMessageID *string          `json:"parent_message_id"`
	Model           string           `json:"model"`
}

type RequestMessage struct {
	ID      string         `json:"id"`
	Role    string         `json:"role"`
	Content RequestContent `json:"content"`
}

type RequestContent struct {
	ContentType string   `json:"content_type"`
	Parts       []string `json:"parts"`
}

func NewRequest(question, conversationID, parentMessageID string) *Request {
	return &Request{
		Action:         "next",
		ConversationID: pointer.ToPointerOrNil(conversationID),
		Messages: []RequestMessage{
			{
				ID:   uuid.New().String(),
				Role: "user",
				Content: RequestContent{
					ContentType: "text",
					Parts:       []string{question},
				},
			},
		},
		ParentMessageID: pointer.ToPointerOrNil(parentMessageID),
		Model:           "text-davinci-002-render",
	}
}
