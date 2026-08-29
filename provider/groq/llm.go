package groq

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a Groq LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM(chat.Compat{
		Name:         "GroqLLM",
		BaseURL:      baseURL,
		DefaultModel: defaultLLMModel,
	}, cfg)
}
