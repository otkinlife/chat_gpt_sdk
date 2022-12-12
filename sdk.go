package chatgpt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-infrastructure/go-ChatGPT/entity"
	"github.com/golang-infrastructure/go-ChatGPT/lib"
	"github.com/google/uuid"
	"golang.org/x/crypto/ssh/terminal"
)

const ConversationAPIURL = "https://chat.openai.com/backend-api/conversation"

type ChatGPT struct {
	token         string
	authorization string

	userAgent string

	conversationID  string
	parentMessageID string
	term            *terminal.Terminal
	Old             string
	Last            string
}

func NewChatGPT(token string, term *terminal.Terminal) *ChatGPT {
	return &ChatGPT{
		token:         token,
		authorization: "Bearer " + token,
		userAgent:     "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36",
		term:          term,
	}
}

func (cg *ChatGPT) Talk(question string) error {
	if cg.parentMessageID == "" {
		cg.parentMessageID = uuid.New().String()
	}
	c, err := cg.sendConversatio(question, cg.conversationID, cg.parentMessageID)
	if err != nil {
		return err
	}
	return cg.getConversatio(c)
}

func (cg *ChatGPT) sendConversatio(question, conversationID, parentMessageID string) (chan string, error) {
	// 设置请求body
	reqData := entity.NewRequest(question, conversationID, parentMessageID)
	reqBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", ConversationAPIURL, bytes.NewBuffer([]byte(reqBytes)))
	if err != nil {
		return nil, err
	}

	// 设置header
	req.Header.Set("User-Agent", cg.userAgent)
	req.Header.Set("Authorization", cg.authorization)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return nil, err
	}
	c := make(chan string, 20)
	go func() {
		err = lib.Post(req, c)
	}()
	return c, err
}

func (cg *ChatGPT) getConversatio(c chan string) error {
	if c == nil {
		return errors.New("get stream error")
	}
	// 创建一个读取器，用于读取响应流
	for {
		v, ok := <-c
		if !ok {
			break
		}
		data := &entity.Response{}
		if len(v) <= 5 {
			continue
		}
		bytelist := []byte(v)[5 : len(v)-1]
		json.Unmarshal(bytelist, data)
		if data.Error != nil {
			return errors.New("get data error")
		}
		lastBytes := []byte(fmt.Sprintf("%s", data.Message.Content.Parts))
		if len(lastBytes) < 2 {
			continue
		}
		cg.Last = string(lastBytes[1 : len(lastBytes)-1])
		d := strings.Replace(cg.Last, cg.Old, "", -1)
		cg.term.Write([]byte(d))
		cg.Old = cg.Last
	}
	cg.term.Write([]byte("\n"))
	return nil
}

func (x *ChatGPT) GetConversationID() string {
	return x.conversationID
}

func (x *ChatGPT) SetConversationID(conversationID string) {
	x.conversationID = conversationID
}

func (x *ChatGPT) GetParentMessageID() string {
	return x.parentMessageID
}

func (x *ChatGPT) SetParentMessageID(parentMessageID string) {
	x.parentMessageID = parentMessageID
}

func (x *ChatGPT) GetUserAgent() string {
	return x.userAgent
}

func (x *ChatGPT) SetUserAgent(userAgent string) {
	x.userAgent = userAgent
}

func (x *ChatGPT) Settoken(token string) {
	x.token = token
	x.authorization = "Bearer " + token
}
