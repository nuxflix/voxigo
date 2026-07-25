package nebius

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds a Nebius AI Studio LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("NebiusLLM", baseURL, defaultModel, cfg)
}
