package ollama

import (
	"net/http"
	"net/url"

	"github.com/ollama/ollama/api"
)

type OllamaClient struct {
	client *api.Client
}

// * NewOllamaClient initializes an Ollama client with your BaseURL and HTTPClient.
func NewOllamaClient(baseURL string, httpClient *http.Client) (*OllamaClient, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	cli := api.NewClient(u, httpClient)
	return &OllamaClient{client: cli}, nil
}
