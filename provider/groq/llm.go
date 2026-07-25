package groq

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds a Groq LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("GroqLLM", baseURL, defaultLLMModel, cfg)
}
