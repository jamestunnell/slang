package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewFloat(val float64) *Expression {
	return NewConst(slang.ExprFLOAT, val)
}
