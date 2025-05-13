package types

import (
	"net/http"
)

type Config struct {
	OllamaConfig    *OllamaConfig
	AnthropicConfig *AnthropicConfig
	// OpenAIConfig    *OpenAIConfig
	// VertexAIConfig  *VertexAIConfig
}

type AnthropicConfig struct {
	APIKey string
}

type OllamaConfig struct {
	BaseURL    string
	HTTPClient *http.Client
}
