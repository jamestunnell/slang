package tokens

import "github.com/jamestunnell/slang"

type Less struct{}

func LESS() slang.TokenInfo           { return &Less{} }
func (t *Less) Type() slang.TokenType { return slang.TokenLESS }
func (t *Less) Value() string         { return "<" }
