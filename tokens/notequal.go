package tokens

import "github.com/jamestunnell/slang"

type NotEqual struct{}

func NOTEQUAL() slang.TokenInfo           { return &NotEqual{} }
func (t *NotEqual) Type() slang.TokenType { return slang.TokenNOTEQUAL }
func (t *NotEqual) Value() string         { return "!=" }
