package xai

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds an xAI (Grok) LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("XAILLM", baseURL, defaultModel, cfg)
}
