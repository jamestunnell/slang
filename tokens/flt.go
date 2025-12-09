package tokens

import "github.com/jamestunnell/slang"

type Flt struct{}

const StrFLOAT = "float"

func FLOAT() slang.TokenInfo         { return &Flt{} }
func (t *Flt) Type() slang.TokenType { return slang.TokenFLOAT }
func (t *Flt) Value() string         { return StrFLOAT }
