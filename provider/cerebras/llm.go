package cerebras

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds a Cerebras LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("CerebrasLLM", baseURL, defaultModel, cfg)
}
