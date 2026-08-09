package novita

import "github.com/nuxflix/voxigo/provider/openai/chat"

// NewLLM builds a Novita AI LLM service.
func NewLLM(cfg chat.LLMConfig) *chat.LLMService {
	return chat.NewCompatLLM(chat.Compat{
		Name:         "NovitaLLM",
		BaseURL:      baseURL,
		DefaultModel: defaultModel,
	}, cfg)
}
