package tokens

import "github.com/jamestunnell/slang"

type Var struct{}

const StrVAR = "var"

func VAR() slang.TokenInfo           { return &Var{} }
func (t *Var) Type() slang.TokenType { return slang.TokenVAR }
func (t *Var) Value() string         { return StrVAR }
