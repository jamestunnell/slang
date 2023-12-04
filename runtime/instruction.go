package runtime

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"
)

type Instruction struct {
	Opcode   Opcode
	Operands []Operand
}

type Instructions []*Instruction

type Operand interface {
	Width() int
	Put([]byte)
}

type Uint8Operand struct {
	Value uint8
}

type Uint16Operand struct {
	Value uint16
}

type Uint32Operand struct {
	Value uint32
}

type Uint64Operand struct {
	Value uint64
}

func NewInstruction(opcode Opcode, operands ...Operand) *Instruction {
	return &Instruction{
		Opcode:   opcode,
		Operands: operands,
	}
}

func NewUint8Operand(val uint8) *Uint8Operand {
	return &Uint8Operand{Value: val}
}

func NewUint16Operand(val uint16) *Uint16Operand {
	return &Uint16Operand{Value: val}
}

func NewUint32Operand(val uint32) *Uint32Operand {
	return &Uint32Operand{Value: val}
}

func NewUint64Operand(val uint64) *Uint64Operand {
	return &Uint64Operand{Value: val}
}

func FormatOperand(operand Operand) string {
	d := make([]byte, operand.Width())

	operand.Put(d)

	return hex.EncodeToString(d)
}

func (i *Instruction) LengthBytes() int {
	length := 1 // for the op code

	for _, o := range i.Operands {
		length += o.Width()
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

			ptr += operand.Width()
		}
	}

	return data
}

func (operand *Uint8Operand) Width() int {
	return 1
}

func (operand *Uint8Operand) Put(d []byte) {
	d[0] = operand.Value
}

func (operand *Uint16Operand) Width() int {
	return 2
}

func (operand *Uint16Operand) Put(d []byte) {
	binary.BigEndian.PutUint16(d, operand.Value)
}

func (operand *Uint32Operand) Width() int {
	return 4
}

func (operand *Uint32Operand) Put(d []byte) {
	binary.BigEndian.PutUint32(d, operand.Value)
}

func (operand *Uint64Operand) Width() int {
	return 8
}

func (operand *Uint64Operand) Put(d []byte) {
	binary.BigEndian.PutUint64(d, operand.Value)
}
