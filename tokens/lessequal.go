package tokens

import "github.com/jamestunnell/slang"

type LessEqual struct{}

func LESSEQUAL() slang.TokenInfo           { return &LessEqual{} }
func (t *LessEqual) Type() slang.TokenType { return slang.TokenLESSEQUAL }
func (t *LessEqual) Value() string         { return "<=" }
