package ollama

import (
	"context"

	"modelgene"
	"modelgene/pkg/types"

	"github.com/ollama/ollama/api"
)

type Provider struct {
	client *OllamaClient
	config *types.OllamaConfig
}

func NewProvider(cfg *types.OllamaConfig) (*Provider, error) {
	if cfg == nil {
		return nil, modelgene.NewError(types.ProviderOllama, "ollama config is nil", nil)
	}

	cli, err := NewOllamaClient(cfg.BaseURL, cfg.HTTPClient)
	if err != nil {
		return nil, modelgene.NewError(types.ProviderOllama, "failed to create client", err)
	}

	return &Provider{
		client: cli,
		config: cfg,
	}, nil
}

func (p *Provider) Chat(ctx context.Context, req types.APIRequest) (*types.APIResponse, error) {
	if req.Model == "" {
		return nil, modelgene.NewError(types.ProviderOllama, "model name is required", nil)
	}

	genReq := &api.ChatRequest{
		Model:    req.Model,
		Messages: convertMessages(req.Messages),
		Stream:   modelgene.PtrBool(false),
		Options:  req.OllamaOptions,
	}

	var fullResponse string
	var finishReason string

	err := p.client.client.Chat(ctx, genReq, func(resp api.ChatResponse) error {
		fullResponse += resp.Message.Content
		finishReason = resp.DoneReason
		return nil
	})
	if err != nil {
		return nil, modelgene.NewError(types.ProviderOllama, "chat error", err)
	}

	response := &types.APIResponse{
		Model:    req.Model,
		Provider: types.ProviderOllama,
		Choices: []types.Choice{
			{
				Index: 0,
				Message: types.Message{
					Role:    "assistant",
					Content: fullResponse,
				},
				FinishReason: finishReason,
			},
		},
	}

	return response, nil
}

func convertMessages(msgs []types.Message) []api.Message {
	var out []api.Message
	for _, m := range msgs {
		out = append(out, api.Message{
			Role:    m.Role,
			Content: m.Content,
		})
	}
	return out
}
