package tokens

import "github.com/jamestunnell/slang"

type MinusEqual struct{}

func MINUSEQUAL() slang.TokenInfo           { return &MinusEqual{} }
func (t *MinusEqual) Type() slang.TokenType { return slang.TokenMINUSEQUAL }
func (t *MinusEqual) Value() string         { return "-=" }
