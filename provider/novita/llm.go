package novita

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds a Novita AI LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("NovitaLLM", baseURL, defaultModel, cfg)
}
