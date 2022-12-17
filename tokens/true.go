package tokens

import "github.com/jamestunnell/slang"

type True struct{}

const StrTRUE = "true"

func TRUE() slang.TokenInfo           { return &True{} }
func (t *True) Type() slang.TokenType { return slang.TokenTRUE }
func (t *True) Value() string         { return StrTRUE }
