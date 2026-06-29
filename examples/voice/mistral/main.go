// Command voice-mistral is the voice bot with a Mistral LLM, served through
// voxigo's OpenAI-compatible client (Deepgram STT, ElevenLabs TTS). Run it, open
// the URL it logs, click start and allow the mic:
//
//	DEEPGRAM_API_KEY=… MISTRAL_API_KEY=… ELEVENLABS_API_KEY=… go run ./examples/voice/mistral
package main

import (
	"log"
	"os"

	"github.com/nuxflix/voxigo/examples/voice/run"
	"github.com/nuxflix/voxigo/provider/mistral"
	"github.com/nuxflix/voxigo/provider/openai"
)

func main() {
	if err := run.Voice(run.Builders{
		STT: run.DefaultSTT,
		LLM: run.Service(openai.LLMConfig{APIKey: os.Getenv("MISTRAL_API_KEY")}, mistral.NewLLM),
		TTS: run.DefaultTTS,
	}); err != nil {
		log.Fatal(err)
	}
}
