package processor_test

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/nuxflix/voxigo/clock"
	"github.com/nuxflix/voxigo/frames"
	"github.com/nuxflix/voxigo/processor"
)

func pcm16(samples ...int16) []byte {
	b := make([]byte, len(samples)*2)
	for i, s := range samples {
		binary.LittleEndian.PutUint16(b[i*2:], uint16(s))
	}
	return b
}

func TestAudioGainScalesInputAndPassesOtherFrames(t *testing.T) {
	g := processor.NewAudioGain(2)
	c := newCapture()
	g.Link(c)

	ctx := context.Background()
	setup := processor.Setup{Clock: clock.NewSystem()}
	if err := g.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	if err := c.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = g.Cleanup(ctx)
		_ = c.Cleanup(ctx)
	})

	_ = g.QueueFrame(ctx, frames.NewStartFrame(), processor.Downstream)
	mustReceive[*frames.StartFrame](t, c.got, "StartFrame")

	in := frames.NewInputAudioRawFrame(pcm16(100, -50), 16000, 1)
	_ = g.QueueFrame(ctx, in, processor.Downstream)
	got := mustReceive[*frames.InputAudioRawFrame](t, c.got, "InputAudioRawFrame")
	s0 := int16(binary.LittleEndian.Uint16(got.Audio[0:]))
	s1 := int16(binary.LittleEndian.Uint16(got.Audio[2:]))
	if s0 != 200 || s1 != -100 {
		t.Fatalf("gained samples = %d, %d, want 200, -100", s0, s1)
	}

	text := frames.NewTextFrame("leave me")
	_ = g.QueueFrame(ctx, text, processor.Downstream)
	out := mustReceive[*frames.TextFrame](t, c.got, "TextFrame")
	if out.Text != "leave me" {
		t.Fatalf("Text = %q", out.Text)
	}
}
