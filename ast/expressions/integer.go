package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewInteger(val int64) *Const[int64] {
	return NewConst(slang.ExprINTEGER, val)
}
