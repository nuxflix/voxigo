package openrouter_test

import (
	"testing"

	"github.com/nuxflix/voxigo/internal/providertest"
	"github.com/nuxflix/voxigo/provider/openrouter"
)

// TestNewLLM checks the OpenRouter shim wires the right service name and
// default model into the shared OpenAI-compatible client.
func TestNewLLM(t *testing.T) {
	providertest.CompatLLM(t, "OpenRouterLLM", "openai/gpt-4.1", openrouter.NewLLM)
}
