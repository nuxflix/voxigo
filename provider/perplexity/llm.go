package perplexity

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a Perplexity LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM(chat.Compat{
		Name:            "PerplexityLLM",
		BaseURL:         baseURL,
		DefaultModel:    defaultModel,
		NoDeveloperRole: true,
	}, cfg)
}
