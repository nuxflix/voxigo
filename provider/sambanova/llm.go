package sambanova

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds a SambaNova Cloud LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("SambaNovaLLM", baseURL, defaultModel, cfg)
}
