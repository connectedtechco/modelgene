package client

import (
	"context"

	"github.com/connectedtechco/modelgene"
	"github.com/connectedtechco/modelgene/pkg/types"
	"github.com/connectedtechco/modelgene/providers/anthropic"
	"github.com/connectedtechco/modelgene/providers/ollama"
)

// * Client is a unified entry point for interacting with multiple provider clients.
type Client struct {
	providers map[types.Provider]types.ProviderClient
}

// * NewClient initializes all available providers from the given config.
// * It returns a Client with a map of active provider clients.
func NewClient(cfg *types.Config) (*Client, error) {
	providers := make(map[types.Provider]types.ProviderClient)

	// * Register Ollama provider if config is present.
	if cfg.OllamaConfig != nil {
		ollamaProvider, err := ollama.NewProvider(cfg.OllamaConfig)
		if err != nil {
			return nil, modelgene.NewError(types.ProviderOllama, "failed to init provider", err)
		}
		providers[types.ProviderOllama] = ollamaProvider
	}

	// * Register Anthropic provider if config is present.
	if cfg.AnthropicConfig != nil {
		anthropicProvider, err := anthropic.NewProvider(cfg.AnthropicConfig)
		if err != nil {
			return nil, modelgene.NewError(types.ProviderAnthropic, "failed to init provider", err)
		}
		providers[types.ProviderAnthropic] = anthropicProvider
	}

	// * Return the fully initialized client.
	return &Client{
		providers: providers,
	}, nil
}

// * Chat routes the chat request to the specified provider's Chat implementation.
func (c *Client) Chat(ctx context.Context, provider types.Provider, req types.APIRequest) (*types.APIResponse, error) {
	prov, ok := c.providers[provider]
	if !ok {
		return nil, modelgene.NewError(provider, "provider is not configured", nil)
	}
	return prov.Chat(ctx, req)
}

// * Embed checks if the provider supports embedding and routes the request.
// * Returns an error if the provider does not implement the Embedder interface.
func (c *Client) Embed(ctx context.Context, provider types.Provider, req types.APIRequest) (*types.APIResponse, error) {
	prov, ok := c.providers[provider]
	if !ok {
		return nil, modelgene.NewError(provider, "provider is not configured", nil)
	}

	embedder, ok := prov.(types.Embedder)
	if !ok {
		return nil, modelgene.NewError(provider, "provider does not support embedding", nil)
	}

	return embedder.Embed(ctx, &req)
}
