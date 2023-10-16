package tokens

import "github.com/jamestunnell/slang"

type Use struct{}

const StrUSE = "use"

func USE() slang.TokenInfo           { return &Use{} }
func (t *Use) Type() slang.TokenType { return slang.TokenUSE }
func (t *Use) Value() string         { return StrUSE }
