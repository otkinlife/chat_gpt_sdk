package completions

// ChatCompletionResponse 定义了整个响应的结构
type ChatCompletionResponse struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int64       `json:"created"`
	Model             string      `json:"model"`
	Choices           []Choice    `json:"choices"`
	Usage             Usage       `json:"usage"`
	SystemFingerprint interface{} `json:"system_fingerprint"`
}

// Choice 定义了 "choices" 数组中的元素结构
type Choice struct {
	Index        int         `json:"index"`
	Message      Message     `json:"message"`
	Delta        Message     `json:"Delta"`
	LogProbs     interface{} `json:"logprobs"`
	FinishReason string      `json:"finish_reason"`
}

// Message 定义了 "message" 对象的结构
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// Usage 定义了 "usage" 对象的结构
type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}
