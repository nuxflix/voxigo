package together

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a Together AI LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM("TogetherLLM", baseURL, defaultModel, cfg)
}
