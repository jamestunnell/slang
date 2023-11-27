package tokens

import (
	"github.com/jamestunnell/slang"
)

type VerbatimString struct{ val string }

func VERBATIMSTRING(val string) slang.TokenInfo { return &VerbatimString{val: val} }
func (t *VerbatimString) Type() slang.TokenType { return slang.TokenVERBATIMSTRING }
func (t *VerbatimString) Value() string         { return t.val }
