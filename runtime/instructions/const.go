package instructions

import "github.com/jamestunnell/slang/runtime"

func NewGetConst(idx uint16) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpGETCONST, runtime.NewUint16Operand(idx))
}
