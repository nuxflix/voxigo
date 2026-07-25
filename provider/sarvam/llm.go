package sarvam

import "github.com/nuxflix/voxigo/provider/openai"

const (
	llmBaseURL      = "https://api.sarvam.ai/v1"
	defaultLLMModel = "sarvam-30b"
)

// NewLLM builds a Sarvam LLM service over its OpenAI-compatible endpoint.
func NewLLM(cfg openai.LLMConfig) *openai.LLMService {
	return openai.NewCompatLLM("SarvamLLM", llmBaseURL, defaultLLMModel, cfg)
}
