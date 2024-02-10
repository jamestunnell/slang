package tokens

import "github.com/jamestunnell/slang"

type Const struct{}

const StrCONST = "const"

func CONST() slang.TokenInfo           { return &Const{} }
func (t *Const) Type() slang.TokenType { return slang.TokenCONST }
func (t *Const) Value() string         { return StrCONST }
