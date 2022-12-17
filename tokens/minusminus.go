package tokens

import "github.com/jamestunnell/slang"

type MinusMinus struct{}

func MINUSMINUS() slang.TokenInfo           { return &MinusMinus{} }
func (t *MinusMinus) Type() slang.TokenType { return slang.TokenMINUSMINUS }
func (t *MinusMinus) Value() string         { return "--" }
