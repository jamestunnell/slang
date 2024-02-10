package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewFloat(val float64) *Const[float64] {
	return NewConst(slang.ExprFLOAT, val)
}
