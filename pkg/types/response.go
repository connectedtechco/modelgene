package types

// * Choice represents a single completion choice within an APIResponse.
type Choice struct {
	Index        int          `json:"index"`
	Message      Message      `json:"message"` // Role should be "assistant"
	FinishReason string       `json:"finish_reason"`
	LogProbs     *LogProbInfo `json:"logprobs,omitempty"`
}

// * APIResponse represents a unified response structure from LLM APIs.
type APIResponse struct {
	ID       string   `json:"id,omitempty"`
	Model    string   `json:"model"`
	Provider Provider `json:"provider"`
	Choices  []Choice `json:"choices"`
	Usage    *Usage   `json:"usage,omitempty"`
}

// * Usage represents token usage statistics.
type Usage struct {
	PromptTokens     *int `json:"prompt_tokens,omitempty"`
	CompletionTokens *int `json:"completion_tokens,omitempty"`
	TotalTokens      *int `json:"total_tokens,omitempty"`
}

// * LogProbInfo represents log probability information (placeholder).
type LogProbInfo struct {
	Content interface{} `json:"content,omitempty"`
}
