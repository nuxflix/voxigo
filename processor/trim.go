package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioTrim drops leading and trailing silence from every audio frame that
// passes through it. A frame that is silent end to end keeps its shape as an
// empty payload rather than being dropped, so timing and frame order stay
// intact while the samples themselves go away.
type AudioTrim struct {
	*Base
}

// NewAudioTrim builds a processor that strips silence from the edges of each
// audio frame.
func NewAudioTrim() *AudioTrim {
	p := &AudioTrim{}
	p.Base = New("AudioTrim", p)
	return p
}

// ProcessFrame trims audio frames and forwards everything else.
func (p *AudioTrim) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := p.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		if trimmed := audio.TrimSilence(data.Audio); trimmed != nil {
			data.Audio = trimmed
		} else {
			data.Audio = data.Audio[:0]
		}
	}
	return p.PushFrame(ctx, f, dir)
}
