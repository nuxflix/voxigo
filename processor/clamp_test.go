package processor_test

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/nuxflix/voxigo/clock"
	"github.com/nuxflix/voxigo/frames"
	"github.com/nuxflix/voxigo/processor"
)

func pcm16clamp(samples ...int16) []byte {
	b := make([]byte, len(samples)*2)
	for i, s := range samples {
		binary.LittleEndian.PutUint16(b[i*2:], uint16(s))
	}
	return b
}

func TestAudioClampLimitsPeaksAndPassesOtherFrames(t *testing.T) {
	p := processor.NewAudioClamp(100)
	c := newCapture()
	p.Link(c)

	ctx := context.Background()
	setup := processor.Setup{Clock: clock.NewSystem()}
	if err := p.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	if err := c.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = p.Cleanup(ctx)
		_ = c.Cleanup(ctx)
	})

	_ = p.QueueFrame(ctx, frames.NewStartFrame(), processor.Downstream)
	mustReceive[*frames.StartFrame](t, c.got, "StartFrame")

	in := frames.NewInputAudioRawFrame(pcm16clamp(200, -200, 10), 16000, 1)
	_ = p.QueueFrame(ctx, in, processor.Downstream)
	got := mustReceive[*frames.InputAudioRawFrame](t, c.got, "InputAudioRawFrame")
	s0 := int16(binary.LittleEndian.Uint16(got.Audio[0:]))
	s1 := int16(binary.LittleEndian.Uint16(got.Audio[2:]))
	s2 := int16(binary.LittleEndian.Uint16(got.Audio[4:]))
	if s0 != 100 || s1 != -100 || s2 != 10 {
		t.Fatalf("clamped samples = %d, %d, %d, want 100, -100, 10", s0, s1, s2)
	}

	text := frames.NewTextFrame("leave me")
	_ = p.QueueFrame(ctx, text, processor.Downstream)
	out := mustReceive[*frames.TextFrame](t, c.got, "TextFrame")
	if out.Text != "leave me" {
		t.Fatalf("Text = %q", out.Text)
	}
}
