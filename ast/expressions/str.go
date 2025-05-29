package expressions

import (
	"github.com/jamestunnell/slang"
)

func NewStr(val string) *Expression {
	return NewConst(slang.ExprSTR, val)
}
