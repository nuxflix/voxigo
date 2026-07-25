package together

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds a Together AI LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("TogetherLLM", baseURL, defaultModel, cfg)
}
