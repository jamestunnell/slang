package tokens

import "github.com/jamestunnell/slang"

type GreaterEqual struct{}

func GREATEREQUAL() slang.TokenInfo           { return &GreaterEqual{} }
func (t *GreaterEqual) Type() slang.TokenType { return slang.TokenGREATEREQUAL }
func (t *GreaterEqual) Value() string         { return ">=" }
