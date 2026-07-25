package kokoro

import (
	"github.com/nuxflix/voxigo/provider/openai"
	"github.com/nuxflix/voxigo/service/tts"
)

// NewTTS builds a Kokoro TTS service. Kokoro-FastAPI returns 24 kHz PCM, which
// matches the OpenAI "pcm" response format this wrapper requests.
func NewTTS(cfg Config) *tts.Base {
	return openai.NewCompatTTS("KokoroTTS", cfg.BaseURL, defaultModel, defaultVoice, openai.TTSConfig{
		APIKey:  cfg.APIKey,
		BaseURL: cfg.BaseURL,
		Model:   cfg.Model,
		Voice:   cfg.Voice,
		Speed:   cfg.Speed,
	})
}
