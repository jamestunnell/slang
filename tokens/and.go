package tokens

import "github.com/jamestunnell/slang"

type And struct{}

const StrAND = "and"

func AND() slang.TokenInfo           { return &And{} }
func (t *And) Type() slang.TokenType { return slang.TokenAND }
func (t *And) Value() string         { return StrAND }
