// Package completions Description: This file contains the code for the chat endpoint of the OpenAI API.
// DOC:https://platform.openai.com/docs/api-reference/chat
package completions

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// OpenAIClient Define a struct for the OpenAI client
type OpenAIClient struct {
	APIKey     string
	BaseURL    string
	HttpClient *http.Client
}

// NewClient Constructor for the OpenAI client
func NewClient(apiKey string) *OpenAIClient {
	return &OpenAIClient{
		APIKey:     apiKey,
		BaseURL:    "https://api.openai.com/v1",
		HttpClient: &http.Client{},
	}
}

// CreateChat Create a chat session
func (client *OpenAIClient) CreateChat(request ChatCompletionRequest) (*ChatCompletionResponse, error) {
	reqBody, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", client.BaseURL+"/chat/completions", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+client.APIKey)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// 读取响应体以获取更多错误信息
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(respBody))
	}
	respBody, err := io.ReadAll(resp.Body)
	fmt.Println(string(respBody))
	var chatResponse ChatCompletionResponse
	if err := json.NewDecoder(resp.Body).Decode(&chatResponse); err != nil {
		return nil, err
	}

	return &chatResponse, nil
}

// CreateChatStream 创建一个处理流式响应的聊天会话
func (client *OpenAIClient) CreateChatStream(request ChatCompletionRequest) (<-chan ChatCompletionResponse, <-chan error) {
	responseChan := make(chan ChatCompletionResponse)
	errorChan := make(chan error, 1) // 缓冲为1，以防止阻塞

	go func() {
		defer close(responseChan)
		defer close(errorChan)

		reqBody, err := json.Marshal(request)
		if err != nil {
			errorChan <- err
			return
		}

		req, err := http.NewRequest("POST", client.BaseURL+"/chat/completions", bytes.NewBuffer(reqBody))
		if err != nil {
			errorChan <- err
			return
		}

		req.Header.Add("Authorization", "Bearer "+client.APIKey)
		req.Header.Add("Content-Type", "application/json")

		resp, err := client.HttpClient.Do(req)
		if err != nil {
			errorChan <- err
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			body, _ := io.ReadAll(resp.Body)
			errorChan <- fmt.Errorf("API request failed with status code %d: %s", resp.StatusCode, string(body))
			return
		}
		scanner := bufio.NewScanner(resp.Body)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.HasPrefix(line, "data: ") {
				var chunk ChatCompletionResponse
				err := json.Unmarshal([]byte(line[6:]), &chunk) // 从 "data: " 后开始解析
				if err != nil {
					errorChan <- err
					return
				}
				responseChan <- chunk
			} else if line == "data: [DONE]" {
				// 流式响应结束
				return
			}
		}
		if err := scanner.Err(); err != nil {
			errorChan <- err
			return
		}
	}()
	return responseChan, errorChan
}
