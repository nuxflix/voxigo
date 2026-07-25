package deepseek_test

import (
	"testing"

	"github.com/nuxflix/voxigo/internal/providertest"
	"github.com/nuxflix/voxigo/provider/deepseek"
)

// TestNewLLM checks the DeepSeek shim wires the right service name and
// default model into the shared OpenAI-compatible client.
func TestNewLLM(t *testing.T) {
	providertest.CompatLLM(t, "DeepSeekLLM", "deepseek-chat", deepseek.NewLLM)
}
