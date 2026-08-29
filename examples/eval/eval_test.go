package main

import (
	"testing"

	"github.com/nuxflix/voxigo/eval"
)

// TestGreeting plays the greeting scenario against the demo bot in-process — the
// same scenario `voxigo eval run` drives over a WebSocket.
func TestGreeting(t *testing.T) {
	eval.Run(t, "scenarios/greeting.yaml", buildBot)
}
