package tokens

import (
	"github.com/jamestunnell/slang"
)

type IntVal struct {
	val string
}

func INTVAL(val string) slang.TokenInfo { return &IntVal{val: val} }
func (t *IntVal) Type() slang.TokenType { return slang.TokenINTVAL }
func (t *IntVal) Value() string         { return t.val }
