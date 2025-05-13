package ollama

import (
	"context"
	"strings"
	"fmt"

	"github.com/connectedtechco/modelgene"
	"github.com/connectedtechco/modelgene/pkg/types"

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

func (p *Provider) Embed(ctx context.Context, req *api.EmbedRequest) (*types.APIResponse, error) {
	if req.Model == "" {
		return nil, modelgene.NewError(types.ProviderOllama, "embedding model name is required", nil)
	}

	resp, err := p.client.client.Embed(ctx, req)
	if err != nil {
		return nil, modelgene.NewError(types.ProviderOllama, "embedding error", err)
	}

	// * Serialize the embedding vector as a string (comma-separated)
	vectorStrings := make([]string, len(resp.Embeddings))
	for i, vec := range resp.Embeddings {
		var parts []string
		for _, v := range vec {
			parts = append(parts, fmt.Sprintf("%f", v))
		}
		vectorStrings[i] = strings.Join(parts, ",")
	}

	// * store the vector string in `Message.Content`
	choices := make([]types.Choice, len(vectorStrings))
	for i, vs := range vectorStrings {
		choices[i] = types.Choice{
			Index: i,
			Message: types.Message{
				Role:    "assistant",
				Content: vs,
			},
			FinishReason: "stop",
		}
	}

	return &types.APIResponse{
		Model:    req.Model,
		Provider: types.ProviderOllama,
		Choices:  choices,
	}, nil
}
