package completions

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestOpenAIClient_CreateChat(t *testing.T) {

	client := NewClient(os.Getenv("APIKEY"))

	request := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		ResponseFormat: ResponseFormat{
			Type: "text",
		},
		Messages: []Message{
			{
				Role:    "user",
				Content: "你好",
			},
		},
		Stream:    false,
		N:         1,
		MaxTokens: 2048,
	}

	response, err := client.CreateChat(request)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Response: %+v\n", response)
}

func TestOpenAIClient_CreateChatStream(t *testing.T) {

	client := NewClient(os.Getenv("APIKEY"))

	request := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		ResponseFormat: ResponseFormat{
			Type: "text",
		},
		Messages: []Message{
			{
				Role:    "user",
				Content: "你好",
			},
		},
		Stream:    true,
		N:         1,
		MaxTokens: 2048,
	}

	responseChan, errorChan := client.CreateChatStream(request)
	for {
		select {
		case response, ok := <-responseChan:
			if !ok {
				// channel 已关闭，流式响应结束
				fmt.Println("Stream ended.")
				return
			}
			fmt.Printf("Received chunk: %+v\n", response)
		case err := <-errorChan:
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
		}
	}
}
