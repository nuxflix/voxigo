package qwen_test

import (
	"testing"

	"github.com/nuxflix/voxigo/internal/providertest"
	"github.com/nuxflix/voxigo/provider/qwen"
)

// TestNewLLM checks the Qwen (DashScope) shim wires the right service name and
// default model into the shared OpenAI-compatible client.
func TestNewLLM(t *testing.T) {
	providertest.CompatLLM(t, "QwenLLM", "qwen-plus", qwen.NewLLM)
}
