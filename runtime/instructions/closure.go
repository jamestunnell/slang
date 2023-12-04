package instructions

import "github.com/jamestunnell/slang/runtime"

func NewClosure(fnIdx uint16, numFreeVars uint8) *runtime.Instruction {
	return runtime.NewInstruction(
		runtime.OpCLOSURE,
		runtime.NewUint16Operand(fnIdx),
		runtime.NewUint8Operand(numFreeVars),
	)
}

func NewCurrentClosure() *runtime.Instruction {
	return runtime.NewInstruction(runtime.OpCURRENTCLOSURE)
}
