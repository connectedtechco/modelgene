package types

import "encoding/json"

// Message represents a single message in the conversation history.
type Message struct {
	Role       string     `json:"role"`
	Content    string     `json:"content"`
	Images     []string   `json:"images,omitempty"`       // Used by Ollama
	Name       *string    `json:"name,omitempty"`         // Used by OpenAI
	ToolCallID *string    `json:"tool_call_id,omitempty"` // Used by OpenAI
	ToolCalls  []ToolCall `json:"tool_calls,omitempty"`   // Used by OpenAI
}

// ToolCall represents a tool invocation requested by the model.
type ToolCall struct {
	ID       string           `json:"id"`
	Type     string           `json:"type"` // Usually "function"
	Function ToolCallFunction `json:"function"`
}

// ToolCallFunction details the function called by the model.
type ToolCallFunction struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"` // JSON string arguments
}

// Tool represents a tool/function definition provided to the model.
type Tool struct {
	Type        string              `json:"type,omitempty"`         // Used by OpenAI ("function")
	Function    *FunctionDefinition `json:"function,omitempty"`     // Used by OpenAI
	Name        *string             `json:"name,omitempty"`         // Used by Anthropic/Vertex AI
	Description *string             `json:"description,omitempty"`  // Used by Anthropic/Vertex AI
	InputSchema json.RawMessage     `json:"input_schema,omitempty"` // Used by Anthropic/Vertex AI
}

// FunctionDefinition describes a function for OpenAI tools.
type FunctionDefinition struct {
	Name        string          `json:"name"`
	Description *string         `json:"description,omitempty"`
	Parameters  json.RawMessage `json:"parameters"` // JSON Schema object
}

// ToolChoice controls how the model uses tools.
type ToolChoice struct {
	Type     string              `json:"type"` // "none", "auto", "any", "tool", "function"
	Name     *string             `json:"name,omitempty"`
	Function *ToolChoiceFunction `json:"function,omitempty"` // OpenAI specific structure
}

// ToolChoiceFunction specifies a function name for OpenAI's tool_choice.
type ToolChoiceFunction struct {
	Name string `json:"name"`
}

// ResponseFormat specifies constraints on the output format.
type ResponseFormat struct {
	Type   string `json:"type,omitempty"`   // OpenAI ("text", "json_object")
	Format string `json:"format,omitempty"` // Ollama ("json")
}

// SafetySetting configures content safety filters (primarily Google Vertex AI).
type SafetySetting struct {
	Category  string `json:"category"`
	Threshold string `json:"threshold"`
}

// APIRequest represents the consolidated parameters for calling various LLM APIs.
type APIRequest struct {
	Model            string                 `json:"model"`
	Messages         []Message              `json:"messages,omitempty"`
	Input            string                 `json:"input,omitempty"` // for Ollama Embed()
	Prompt           *string                `json:"prompt,omitempty"`  // Ollama generate endpoint
	System           *string                `json:"system,omitempty"`  // System prompt
	Context          *string                `json:"context,omitempty"` // Vertex AI context
	MaxTokens        *int                   `json:"max_tokens,omitempty"`
	Temperature      *float64               `json:"temperature,omitempty"`
	TopP             *float64               `json:"top_p,omitempty"`
	TopK             *int                   `json:"top_k,omitempty"`
	StopSequences    []string               `json:"stop_sequences,omitempty"` // Renamed from 'stop' for clarity
	PresencePenalty  *float64               `json:"presence_penalty,omitempty"`
	FrequencyPenalty *float64               `json:"frequency_penalty,omitempty"`
	Seed             *int                   `json:"seed,omitempty"`
	Stream           *bool                  `json:"stream,omitempty"` // Although not used by Ollama now, keep for unified API
	N                *int                   `json:"n,omitempty"`      // OpenAI/Vertex AI candidate count
	ResponseFormat   *ResponseFormat        `json:"response_format,omitempty"`
	LogProbs         *bool                  `json:"logprobs,omitempty"`
	TopLogProbs      *int                   `json:"top_logprobs,omitempty"` // OpenAI
	Tools            []Tool                 `json:"tools,omitempty"`
	ToolChoice       *ToolChoice            `json:"tool_choice,omitempty"`
	LogitBias        map[string]int         `json:"logit_bias,omitempty"`      // OpenAI
	User             *string                `json:"user,omitempty"`            // OpenAI
	Metadata         map[string]interface{} `json:"metadata,omitempty"`        // Anthropic
	SafetySettings   []SafetySetting        `json:"safety_settings,omitempty"` // Vertex AI
	Echo             *bool                  `json:"echo,omitempty"`            // Vertex AI
	Template         *string                `json:"template,omitempty"`        // Ollama
	Raw              *bool                  `json:"raw,omitempty"`             // Ollama
	KeepAlive        *string                `json:"keep_alive,omitempty"`      // Ollama
	OllamaOptions    map[string]interface{} `json:"ollama_options,omitempty"`  // Ollama catch-all
}
