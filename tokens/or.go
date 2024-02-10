package tokens

import "github.com/jamestunnell/slang"

type Or struct{}

const StrOR = "or"

func OR() slang.TokenInfo           { return &Or{} }
func (t *Or) Type() slang.TokenType { return slang.TokenOR }
func (t *Or) Value() string         { return StrOR }
