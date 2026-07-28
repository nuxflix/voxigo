package deepseek

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a DeepSeek LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM("DeepSeekLLM", baseURL, defaultModel, cfg)
}
