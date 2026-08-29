package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioInvert flips the polarity of every audio frame that passes through it.
// It is for a session whose microphone or speaker is wired out of phase with
// the rest of the mix, so the samples themselves are inverted rather than the
// hardware.
type AudioInvert struct {
	*Base
}

// NewAudioInvert builds a processor that inverts every audio frame.
func NewAudioInvert() *AudioInvert {
	p := &AudioInvert{}
	p.Base = New("AudioInvert", p)
	return p
}

// ProcessFrame inverts audio frames and forwards everything else.
func (p *AudioInvert) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := p.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		data.Audio = audio.Invert(data.Audio)
	}
	return p.PushFrame(ctx, f, dir)
}
