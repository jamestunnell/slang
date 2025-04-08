package tokens

import "github.com/jamestunnell/slang"

type Interface struct{}

const StrINTERFACE = "interface"

func INTERFACE() slang.TokenInfo           { return &Interface{} }
func (t *Interface) Type() slang.TokenType { return slang.TokenINTERFACE }
func (t *Interface) Value() string         { return StrINTERFACE }
