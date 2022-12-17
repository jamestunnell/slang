package tokens

import "github.com/jamestunnell/slang"

type RBrace struct{}

func RBRACE() slang.TokenInfo           { return &RBrace{} }
func (t *RBrace) Type() slang.TokenType { return slang.TokenRBRACE }
func (t *RBrace) Value() string         { return "}" }
