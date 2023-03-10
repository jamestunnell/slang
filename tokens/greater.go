package tokens

import "github.com/jamestunnell/slang"

type Greater struct{}

func GREATER() slang.TokenInfo           { return &Greater{} }
func (t *Greater) Type() slang.TokenType { return slang.TokenGREATER }
func (t *Greater) Value() string         { return ">" }
