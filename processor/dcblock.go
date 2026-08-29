package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioDCBlock subtracts the mean sample value from every audio frame that
// passes through it. It is for a microphone that sits off zero, so later energy
// measures and VAD see the speech rather than a constant bias.
type AudioDCBlock struct {
	*Base
}

// NewAudioDCBlock builds a processor that removes DC offset from each audio
// frame.
func NewAudioDCBlock() *AudioDCBlock {
	p := &AudioDCBlock{}
	p.Base = New("AudioDCBlock", p)
	return p
}

// ProcessFrame recenters audio frames and forwards everything else.
func (p *AudioDCBlock) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := p.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		data.Audio = audio.RemoveDCOffset(data.Audio)
	}
	return p.PushFrame(ctx, f, dir)
}
