package types

import (
	"context"
	"net/http" // Needed for OllamaConfig
)

// Provider enumerates the supported LLM providers.
type Provider string

const (
	ProviderOllama    Provider = "ollama"
	ProviderOpenAI    Provider = "openai"
	ProviderAnthropic Provider = "anthropic"
	ProviderVertexAI  Provider = "vertexai"
)

// ProviderClient defines the interface that all provider clients must implement.
// It uses APIRequest and APIResponse defined in this package.
type ProviderClient interface {
	Chat(ctx context.Context, req APIRequest) (*APIResponse, error)
}

// Usage represents token usage statistics.
type Usage struct {
	PromptTokens     *int `json:"prompt_tokens,omitempty"`
	CompletionTokens *int `json:"completion_tokens,omitempty"`
	TotalTokens      *int `json:"total_tokens,omitempty"`
}

// LogProbInfo represents log probability information (placeholder).
type LogProbInfo struct {
	Content interface{} `json:"content,omitempty"`
}

type OllamaConfig struct {
	BaseURL    string
	HTTPClient *http.Client
}

