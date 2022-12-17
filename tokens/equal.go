package tokens

import "github.com/jamestunnell/slang"

type Equal struct{}

func EQUAL() slang.TokenInfo           { return &Equal{} }
func (t *Equal) Type() slang.TokenType { return slang.TokenEQUAL }
func (t *Equal) Value() string         { return "==" }
