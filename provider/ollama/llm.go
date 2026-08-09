package ollama

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds an LLM service backed by a local Ollama server. Set cfg.BaseURL
// to reach a remote Ollama; no API key is required by default.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM(chat.Compat{
		Name:            "OllamaLLM",
		BaseURL:         baseURL,
		DefaultModel:    defaultModel,
		NoDeveloperRole: true,
	}, cfg)
}
