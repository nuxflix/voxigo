package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioGain scales the PCM on every audio frame that passes through it. A gain
// of 1 is a no-op. It is for a session that needs the incoming microphone or
// the outgoing speech a little louder or quieter without changing the transport.
type AudioGain struct {
	*Base
	gain float64
}

// NewAudioGain builds a processor that multiplies every audio frame by gain.
func NewAudioGain(gain float64) *AudioGain {
	g := &AudioGain{gain: gain}
	g.Base = New("AudioGain", g)
	return g
}

// ProcessFrame applies the gain to audio frames and forwards everything else.
func (g *AudioGain) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := g.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		data.Audio = audio.ApplyGain(data.Audio, g.gain)
	}
	return g.PushFrame(ctx, f, dir)
}
