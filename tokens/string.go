package tokens

import (
	"github.com/jamestunnell/slang"
)

type String struct{ val string }

func STRING(val string) slang.TokenInfo { return &String{val: val} }
func (t *String) Type() slang.TokenType { return slang.TokenSTRING }
func (t *String) Value() string         { return t.val }
