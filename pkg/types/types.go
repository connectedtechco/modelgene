package types

import (
	"context"
)

// Provider enumerates the supported LLM providers.
type Provider string

const (
	ProviderOllama    Provider = "ollama"
	ProviderOpenAI    Provider = "openai"
	ProviderAnthropic Provider = "anthropic"
	ProviderVertexAI  Provider = "vertexai"
)

// * ProviderClient defines the interface that all provider clients must implement.
type ProviderClient interface {
	Chat(ctx context.Context, req APIRequest) (*APIResponse, error)
}

// * Embedder defines the interface that ollama clients must implement to call Emded().
type Embedder interface {
	Embed(ctx context.Context, req *APIRequest) (*APIResponse, error)
}
