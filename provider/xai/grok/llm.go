package grok

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds an xAI (Grok) LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM(chat.Compat{
		Name:         "XAILLM",
		BaseURL:      baseURL,
		DefaultModel: defaultModel,
	}, cfg)
}
