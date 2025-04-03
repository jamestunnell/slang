package tokens

import "github.com/jamestunnell/slang"

type Str struct{}

const StrSTR = "str"

func STR() slang.TokenInfo           { return &Str{} }
func (t *Str) Type() slang.TokenType { return slang.TokenSTR }
func (t *Str) Value() string         { return StrSTR }
