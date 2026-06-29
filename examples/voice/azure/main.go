// Command voice-azure is the voice bot with Azure for speech-to-text (Azure
// OpenAI Whisper) and text-to-speech (Azure Speech), with an Anthropic LLM. Run
// it, open the URL it logs, click start and allow the mic:
//
//	AZURE_OPENAI_ENDPOINT=… AZURE_OPENAI_API_KEY=… AZURE_STT_DEPLOYMENT=… \
//	AZURE_SPEECH_KEY=… AZURE_SPEECH_REGION=… ANTHROPIC_API_KEY=… \
//	go run ./examples/voice/azure
package main

import (
	"log"
	"os"

	"github.com/nuxflix/voxigo/examples/voice/run"
	"github.com/nuxflix/voxigo/provider/azureopenai"
	"github.com/nuxflix/voxigo/provider/azurespeech"
	"github.com/nuxflix/voxigo/provider/openai"
)

func main() {
	if err := run.Voice(run.Builders{
		STT: run.Service(azureopenai.STTConfig{
			Endpoint:   os.Getenv("AZURE_OPENAI_ENDPOINT"),
			Deployment: os.Getenv("AZURE_STT_DEPLOYMENT"),
			STTConfig:  openai.STTConfig{APIKey: os.Getenv("AZURE_OPENAI_API_KEY"), SampleRate: run.SampleRate},
		}, azureopenai.NewSTT),
		LLM: run.DefaultLLM,
		TTS: run.Service(azurespeech.TTSConfig{
			APIKey: os.Getenv("AZURE_SPEECH_KEY"),
			Region: os.Getenv("AZURE_SPEECH_REGION"),
			Voice:  os.Getenv("AZURE_SPEECH_VOICE"),
		}, azurespeech.NewTTS),
	}); err != nil {
		log.Fatal(err)
	}
}
