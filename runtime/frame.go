package runtime

import "github.com/jamestunnell/slang/runtime/objects"

type Frame struct {
	Closure *objects.Closure

	InstrOffset, InstrLength uint64
	BaseStackCount           int
}

func NewFrame(closure *objects.Closure, baseStackCount int) *Frame {
	return &Frame{
		Closure:        closure,
		InstrOffset:    uint64(0),
		InstrLength:    uint64(len(closure.Func.Instructions)),
		BaseStackCount: baseStackCount,
	}
}

func (f *Frame) Instructions() []byte {
	return f.Closure.Func.Instructions
}
