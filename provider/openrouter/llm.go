package openrouter

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds an OpenRouter LLM service. Set cfg.Model to any "vendor/model"
// slug OpenRouter exposes.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM("OpenRouterLLM", baseURL, defaultModel, cfg)
}
