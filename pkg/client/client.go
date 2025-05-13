package client

import (
	"context"

	"modelgene"
	"modelgene/pkg/types"
	"modelgene/providers/anthropic"
	"modelgene/providers/ollama"
)

type Client struct {
	providers map[types.Provider]types.ProviderClient
}

// NewClient initializes the modelgene client based on available configs
func NewClient(cfg *types.Config) (*Client, error) {
	providers := make(map[types.Provider]types.ProviderClient)

	if cfg.OllamaConfig != nil {
		ollamaProvider, err := ollama.NewProvider(cfg.OllamaConfig)
		if err != nil {
			return nil, modelgene.NewError(types.ProviderOllama, "failed to init provider", err)
		}
		providers[types.ProviderOllama] = ollamaProvider
	}

	if cfg.AnthropicConfig != nil {
		anthropicProvider, err := anthropic.NewProvider(cfg.AnthropicConfig)
		if err != nil {
			return nil, modelgene.NewError(types.ProviderAnthropic, "failed to init provider", err)
		}
		providers[types.ProviderAnthropic] = anthropicProvider
	}

	return &Client{
		providers: providers,
	}, nil
}

func (c *Client) Chat(ctx context.Context, provider types.Provider, req types.APIRequest) (*types.APIResponse, error) {
	prov, ok := c.providers[provider]
	if !ok {
		return nil, modelgene.NewError(provider, "provider is not configured", nil)
	}
	return prov.Chat(ctx, req)
}
