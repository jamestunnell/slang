package tokens

import "github.com/jamestunnell/slang"

type Struct struct{}

const StrSTRUCT = "struct"

func STRUCT() slang.TokenInfo           { return &Struct{} }
func (t *Struct) Type() slang.TokenType { return slang.TokenSTRUCT }
func (t *Struct) Value() string         { return StrSTRUCT }
