package processor

import (
	"context"

	"github.com/nuxflix/voxigo/audio"
	"github.com/nuxflix/voxigo/frames"
)

// AudioNormalize raises or lowers each audio frame so its peak matches Target.
// A target of 0 uses full scale (32767). Silent frames are left alone.
type AudioNormalize struct {
	*Base
	target int
}

// NewAudioNormalize builds a processor that peak-normalizes every audio frame.
func NewAudioNormalize(target int) *AudioNormalize {
	n := &AudioNormalize{target: target}
	n.Base = New("AudioNormalize", n)
	return n
}

// ProcessFrame normalizes audio frames and forwards everything else.
func (n *AudioNormalize) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := n.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if af, ok := f.(frames.AudioFrame); ok {
		data := af.AudioData()
		data.Audio = audio.PeakNormalize(data.Audio, n.target)
	}
	return n.PushFrame(ctx, f, dir)
}
