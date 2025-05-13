package types

// Choice represents a single completion choice within an APIResponse.
// It uses the Message type defined in this package.
type Choice struct {
	Index        int          `json:"index"`
	Message      Message      `json:"message"` // Role should be "assistant"
	FinishReason string       `json:"finish_reason"`
	LogProbs     *LogProbInfo `json:"logprobs,omitempty"` // LogProbInfo defined in this package
}

// APIResponse represents a unified response structure from LLM APIs.
// It uses Choice and Usage types defined in this package.
type APIResponse struct {
	ID       string   `json:"id,omitempty"`
	Model    string   `json:"model"`
	Provider Provider `json:"provider"` // Provider enum defined in this package
	Choices  []Choice `json:"choices"`
	Usage    *Usage   `json:"usage,omitempty"`
}
