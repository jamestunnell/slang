package runtime

import (
	"encoding/binary"

	"github.com/jamestunnell/slang"
)

type Bytecode struct {
	Constants    []slang.Object
	Instructions []byte
}

const MaxVMConstants = 65535

func NewBytecode() *Bytecode {
	return &Bytecode{
		Constants:    []slang.Object{},
		Instructions: []byte{},
	}
}

func (bc *Bytecode) AddConstant(obj slang.Object) (uint16, bool) {
	if len(bc.Constants) == MaxVMConstants {
		return 0, false
	}

	idx := len(bc.Constants)

	bc.Constants = append(bc.Constants, obj)

	return uint16(idx), true
}

type FixupFunc func(targetIdx uint64)

const DummyTargetIdx = uint64(0xFEEDFEEDFEEDFEED)

func (bc *Bytecode) AddSetGlobal(idx uint16) {
	bc.AddInstructionUint16Operands(OpSETGLOBAL, idx)
}

func (bc *Bytecode) AddGetGlobal(idx uint16) {
	bc.AddInstructionUint16Operands(OpGETGLOBAL, idx)
}

func (bc *Bytecode) AddJumpIfFalse() FixupFunc {
	return bc.addJump(OpJUMPIFFALSE)
}

func (bc *Bytecode) AddJump() FixupFunc {
	return bc.addJump(OpJUMP)
}

func (bc *Bytecode) addJump(opcode Opcode) FixupFunc {
	data := make([]byte, 9)

	data[0] = byte(opcode)

	binary.BigEndian.PutUint64(data[1:], DummyTargetIdx)

	// this is where target index will end up in instructions
	targetIdxLoc := len(bc.Instructions) + 1

	bc.Instructions = append(bc.Instructions, data...)

	return func(targetIdx uint64) {
		binary.BigEndian.PutUint64(bc.Instructions[targetIdxLoc:], targetIdx)
	}
}

func (bc *Bytecode) FixupTargetIndex(loc int, targetIdx uint64) {

}

func (bc *Bytecode) AddInstructionNoOperands(opcode Opcode) {
	bc.Instructions = append(bc.Instructions, byte(opcode))
}

func (bc *Bytecode) AddInstructionUint16Operands(opcode Opcode, operands ...uint16) {
	data := make([]byte, 1+2*len(operands))

	data[0] = byte(opcode)

	for i, operand := range operands {
		binary.BigEndian.PutUint16(data[1+2*i:], operand)
	}

	bc.Instructions = append(bc.Instructions, data...)
}
