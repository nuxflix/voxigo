package nvidia

import "github.com/nuxflix/voxigo/provider/openai"

// NewLLM builds an NVIDIA NIM LLM service.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("NvidiaLLM", baseURL, defaultModel, cfg)
}
