package tokens

import "github.com/jamestunnell/slang"

type FltVal struct {
	val string
}

func FLTVAL(val string) slang.TokenInfo { return &FltVal{val: val} }
func (t *FltVal) Type() slang.TokenType { return slang.TokenFLTVAL }
func (t *FltVal) Value() string         { return t.val }
