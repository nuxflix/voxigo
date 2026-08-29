package processor_test

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/nuxflix/voxigo/clock"
	"github.com/nuxflix/voxigo/frames"
	"github.com/nuxflix/voxigo/processor"
)

func TestAudioMuteSilencesWhenArmed(t *testing.T) {
	m := processor.NewAudioMute()
	if m.Muted() {
		t.Fatal("starts muted")
	}
	c := newCapture()
	m.Link(c)

	ctx := context.Background()
	setup := processor.Setup{Clock: clock.NewSystem()}
	if err := m.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	if err := c.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = m.Cleanup(ctx)
		_ = c.Cleanup(ctx)
	})

	_ = m.QueueFrame(ctx, frames.NewStartFrame(), processor.Downstream)
	mustReceive[*frames.StartFrame](t, c.got, "StartFrame")

	open := frames.NewInputAudioRawFrame(pcm16le(100, -50), 16000, 1)
	_ = m.QueueFrame(ctx, open, processor.Downstream)
	live := mustReceive[*frames.InputAudioRawFrame](t, c.got, "open audio")
	if int16(binary.LittleEndian.Uint16(live.Audio)) != 100 {
		t.Fatal("unmuted audio was changed")
	}

	m.SetMuted(true)
	if !m.Muted() {
		t.Fatal("SetMuted(true) did not stick")
	}
	held := frames.NewInputAudioRawFrame(pcm16le(100, -50), 16000, 1)
	_ = m.QueueFrame(ctx, held, processor.Downstream)
	quiet := mustReceive[*frames.InputAudioRawFrame](t, c.got, "muted audio")
	if int16(binary.LittleEndian.Uint16(quiet.Audio)) != 0 || int16(binary.LittleEndian.Uint16(quiet.Audio[2:])) != 0 {
		t.Fatal("muted audio was not silence")
	}
	if len(quiet.Audio) != len(held.Audio) {
		t.Fatal("mute changed the frame length")
	}
}
