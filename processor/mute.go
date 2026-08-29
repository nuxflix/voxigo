package processor

import (
	"context"
	"sync/atomic"

	"github.com/nuxflix/voxigo/frames"
)

// AudioMute replaces audio with silence while it is muted, keeping the frame
// length so the rest of the pipeline still sees a regular stream. Other frames
// pass through. Mute can be flipped from another goroutine.
type AudioMute struct {
	*Base
	muted atomic.Bool
}

// NewAudioMute builds a mute processor. It starts unmuted.
func NewAudioMute() *AudioMute {
	m := &AudioMute{}
	m.Base = New("AudioMute", m)
	return m
}

// SetMuted turns silence on or off for later audio frames.
func (m *AudioMute) SetMuted(muted bool) {
	m.muted.Store(muted)
}

// Muted reports whether later audio frames are silenced.
func (m *AudioMute) Muted() bool {
	return m.muted.Load()
}

// ProcessFrame zeros audio while muted and forwards everything else.
func (m *AudioMute) ProcessFrame(ctx context.Context, f frames.Frame, dir Direction) error {
	if err := m.Base.ProcessFrame(ctx, f, dir); err != nil {
		return err
	}
	if m.muted.Load() {
		if af, ok := f.(frames.AudioFrame); ok {
			data := af.AudioData()
			if len(data.Audio) > 0 {
				data.Audio = make([]byte, len(data.Audio))
			}
		}
	}
	return m.PushFrame(ctx, f, dir)
}
