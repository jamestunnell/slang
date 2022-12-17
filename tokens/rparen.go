package tokens

import "github.com/jamestunnell/slang"

type RParen struct{}

func RPAREN() slang.TokenInfo           { return &RParen{} }
func (t *RParen) Type() slang.TokenType { return slang.TokenRPAREN }
func (t *RParen) Value() string         { return ")" }
