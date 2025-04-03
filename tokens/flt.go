package tokens

import "github.com/jamestunnell/slang"

type Flt struct{}

const StrFLT = "flt"

func FLT() slang.TokenInfo           { return &Flt{} }
func (t *Flt) Type() slang.TokenType { return slang.TokenFLT }
func (t *Flt) Value() string         { return StrFLT }
