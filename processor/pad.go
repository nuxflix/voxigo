package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioPad extends every audio frame with silence until it holds at least
// minSamples samples. A frame that is already that long is left alone. It is
// for a codec or analyzer that wants a fixed minimum chunk and would rather
// see trailing zeros than a short leftover.
type AudioPad struct {
	*Base
	minSamples int
}

// NewAudioPad builds a processor that pads each audio frame to at least
// minSamples samples.
func NewAudioPad(minSamples int) *AudioPad {
	p := &AudioPad{minSamples: minSamples}
	p.Base = New("AudioPad", p)
	return p
}

// ProcessFrame pads audio frames and forwards everything else.
func (p *AudioPad) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := p.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		data.Audio = audio.PadSilence(data.Audio, p.minSamples)
	}
	return p.PushFrame(ctx, f, dir)
}
