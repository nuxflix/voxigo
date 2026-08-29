package processor_test

import (
	"context"
	"encoding/binary"
	"testing"

	"github.com/nuxflix/voxigo/clock"
	"github.com/nuxflix/voxigo/frames"
	"github.com/nuxflix/voxigo/processor"
)

func pcm16le(samples ...int16) []byte {
	b := make([]byte, len(samples)*2)
	for i, s := range samples {
		binary.LittleEndian.PutUint16(b[i*2:], uint16(s))
	}
	return b
}

func TestAudioNormalizeRaisesPeak(t *testing.T) {
	n := processor.NewAudioNormalize(4000)
	c := newCapture()
	n.Link(c)

	ctx := context.Background()
	setup := processor.Setup{Clock: clock.NewSystem()}
	if err := n.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	if err := c.Setup(ctx, setup); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_ = n.Cleanup(ctx)
		_ = c.Cleanup(ctx)
	})

	_ = n.QueueFrame(ctx, frames.NewStartFrame(), processor.Downstream)
	mustReceive[*frames.StartFrame](t, c.got, "StartFrame")

	in := frames.NewInputAudioRawFrame(pcm16le(1000, -2000), 16000, 1)
	_ = n.QueueFrame(ctx, in, processor.Downstream)
	got := mustReceive[*frames.InputAudioRawFrame](t, c.got, "InputAudioRawFrame")
	s0 := int16(binary.LittleEndian.Uint16(got.Audio[0:]))
	s1 := int16(binary.LittleEndian.Uint16(got.Audio[2:]))
	if s0 != 2000 || s1 != -4000 {
		t.Fatalf("normalized samples = %d, %d, want 2000, -4000", s0, s1)
	}
}
