package tokens

import (
	"github.com/jamestunnell/slang"
)

type StrVal struct{ val string }

func STRVAL(val string) slang.TokenInfo { return &StrVal{val: val} }
func (t *StrVal) Type() slang.TokenType { return slang.TokenSTRVAL }
func (t *StrVal) Value() string         { return t.val }
