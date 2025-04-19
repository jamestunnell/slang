package tokens

import "github.com/jamestunnell/slang"

type Bang struct{}

const StrBANG = "!"

func BANG() slang.TokenInfo           { return &Bang{} }
func (t *Bang) Type() slang.TokenType { return slang.TokenBANG }
func (t *Bang) Value() string         { return StrBANG }
