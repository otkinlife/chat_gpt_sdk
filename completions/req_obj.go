package completions

// ChatCompletionRequest Struct to hold chat request parameters
type ChatCompletionRequest struct {
	Model          string         `json:"model"`
	ResponseFormat ResponseFormat `json:"response_format"`
	Messages       []Message      `json:"messages"`
	Stream         bool           `json:"stream"`
	MaxTokens      int            `json:"max_tokens"`
	N              int            `json:"n"`
	Temperature    float64        `json:"temperature"`
}

type ResponseFormat struct {
	Type string `json:"type"`
}
