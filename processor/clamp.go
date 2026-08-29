package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioClamp limits every sample on every audio frame to ±maxAbs. It is a
// hard limiter: peaks above the cap are flattened rather than scaled, so a
// loud burst does not drag the rest of the utterance down with it.
type AudioClamp struct {
	*Base
	maxAbs int
}

// NewAudioClamp builds a processor that limits each audio frame to ±maxAbs.
func NewAudioClamp(maxAbs int) *AudioClamp {
	p := &AudioClamp{maxAbs: maxAbs}
	p.Base = New("AudioClamp", p)
	return p
}

// ProcessFrame clamps audio frames and forwards everything else.
func (p *AudioClamp) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := p.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		data.Audio = audio.Clamp(data.Audio, p.maxAbs)
	}
	return p.PushFrame(ctx, f, dir)
}
