package fireworks

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a Fireworks AI LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM(chat.Compat{
		Name:         "FireworksLLM",
		BaseURL:      baseURL,
		DefaultModel: defaultModel,
	}, cfg)
}
