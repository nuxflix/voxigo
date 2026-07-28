package mistral

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a Mistral AI LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM("MistralLLM", baseURL, defaultModel, cfg)
}
