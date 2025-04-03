package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewInt(val int64) *Const[int64] {
	return NewConst(slang.ExprINT, val)
}
