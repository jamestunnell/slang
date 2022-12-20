package tokens

import "github.com/jamestunnell/slang"

type LBracket struct{}

func LBRACKET() slang.TokenInfo           { return &LBracket{} }
func (t *LBracket) Type() slang.TokenType { return slang.TokenLBRACKET }
func (t *LBracket) Value() string         { return "[" }
