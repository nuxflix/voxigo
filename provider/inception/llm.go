package inception

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds an Inception LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM("InceptionLLM", baseURL, defaultModel, cfg)
}
