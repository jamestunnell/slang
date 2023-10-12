package tokens

import "github.com/jamestunnell/slang"

type Symbol struct {
	val string
}

func SYMBOL(val string) slang.TokenInfo { return &Symbol{val: val} }
func (t *Symbol) Type() slang.TokenType { return slang.TokenSYMBOL }
func (t *Symbol) Value() string         { return t.val }
