package tokens

import "github.com/jamestunnell/slang"

type RBracket struct{}

func RBRACKET() slang.TokenInfo           { return &RBracket{} }
func (t *RBracket) Type() slang.TokenType { return slang.TokenRBRACKET }
func (t *RBracket) Value() string         { return "]" }
