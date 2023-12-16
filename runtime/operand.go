package runtime

import (
	"encoding/binary"
	"encoding/hex"

	"github.com/jamestunnell/slang"
)

type Operand interface {
	GetSymbol() *slang.Symbol
	GetWidth() int

	Put([]byte)
}

type OperandBase struct {
	symbol *slang.Symbol
	width  int
}

type Uint8Operand struct {
	*OperandBase

	Value uint8
}

type Uint16Operand struct {
	*OperandBase

	Value uint16
}

type Uint32Operand struct {
	*OperandBase

	Value uint32
}

type Uint64Operand struct {
	*OperandBase

	Value uint64
}

func NewUint8Operand(val uint8) *Uint8Operand {
	return &Uint8Operand{
		OperandBase: &OperandBase{symbol: nil, width: 1},
		Value:       val,
	}
}

func NewUint8OperandWithSymbol(val uint8, symbol *slang.Symbol) *Uint8Operand {
	return &Uint8Operand{
		OperandBase: &OperandBase{symbol: symbol, width: 1},
		Value:       val,
	}
}

func NewUint16Operand(val uint16) *Uint16Operand {
	return &Uint16Operand{
		OperandBase: &OperandBase{symbol: nil, width: 2},
		Value:       val,
	}
}

func NewUint16OperandWithSymbol(val uint16, symbol *slang.Symbol) *Uint16Operand {
	return &Uint16Operand{
		OperandBase: &OperandBase{symbol: symbol, width: 2},
		Value:       val,
	}
}

func NewUint32Operand(val uint32) *Uint32Operand {
	return &Uint32Operand{
		Value:       val,
		OperandBase: &OperandBase{symbol: nil, width: 4},
	}
}

func NewUint32OperandWithSymbol(val uint32, symbol *slang.Symbol) *Uint32Operand {
	return &Uint32Operand{
		Value:       val,
		OperandBase: &OperandBase{symbol: symbol, width: 4},
	}
}

func NewUint64Operand(val uint64) *Uint64Operand {
	return &Uint64Operand{
		OperandBase: &OperandBase{symbol: nil, width: 8},
		Value:       val,
	}
}

func NewUint64OperandWithSymbol(val uint64, symbol *slang.Symbol) *Uint64Operand {
	return &Uint64Operand{
		OperandBase: &OperandBase{symbol: symbol, width: 8},
		Value:       val,
	}
}

func FormatOperand(operand Operand) string {
	d := make([]byte, operand.GetWidth())

	operand.Put(d)

	return hex.EncodeToString(d)
}

func (operand *OperandBase) GetSymbol() *slang.Symbol {
	return operand.symbol
}

func (operand *OperandBase) GetWidth() int {
	return operand.width
}

func (operand *Uint8Operand) Put(d []byte) {
	d[0] = operand.Value
}

func (operand *Uint16Operand) Put(d []byte) {
	binary.BigEndian.PutUint16(d, operand.Value)
}

func (operand *Uint32Operand) Put(d []byte) {
	binary.BigEndian.PutUint32(d, operand.Value)
}

func (operand *Uint64Operand) Put(d []byte) {
	binary.BigEndian.PutUint64(d, operand.Value)
}
