package runtime

import (
	"fmt"
	"strings"
)

type Instruction struct {
	Opcode   Opcode
	Operands []Operand
}

type Instructions []*Instruction

func NewInstruction(opcode Opcode, operands ...Operand) *Instruction {
	return &Instruction{
		Opcode:   opcode,
		Operands: operands,
	}
}

func (i *Instruction) LengthBytes() int {
	length := 1 // for the op code

	for _, o := range i.Operands {
		length += o.GetWidth()
	}

	return length
}

func (i *Instruction) String() string {
	if len(i.Operands) == 0 {
		return i.Opcode.String()
	}

	operandStrings := make([]string, len(i.Operands))

	for idx, o := range i.Operands {
		operandStrings[idx] = "0x" + FormatOperand(o)
	}

	operandsStr := strings.Join(operandStrings, ", ")

	return fmt.Sprintf("%s: %s", i.Opcode, operandsStr)
}

func (is Instructions) LengthBytes() int {
	length := 0

	for _, instr := range is {
		length += instr.LengthBytes()
	}

	return length
}

func (is Instructions) Assemble() []byte {
	data := make([]byte, is.LengthBytes())
	ptr := 0

	for _, instr := range is {
		data[ptr] = byte(instr.Opcode)

		ptr++

		for _, operand := range instr.Operands {
			operand.Put(data[ptr:])

			ptr += operand.GetWidth()
		}
	}

	return data
}
