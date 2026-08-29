package nebius_test

import (
	"testing"

	"github.com/nuxflix/voxigo/internal/providertest"
	"github.com/nuxflix/voxigo/provider/nebius"
)

// TestNewLLM checks the Nebius AI Studio shim wires the right service name and
// default model into the shared OpenAI-compatible client.
func TestNewLLM(t *testing.T) {
	providertest.CompatLLM(t, "NebiusLLM", "meta-llama/Llama-3.3-70B-Instruct", nebius.NewLLM)
}
