package tokens

import "github.com/jamestunnell/slang"

type Bool struct{}

const StrBOOL = "bool"

func BOOL() slang.TokenInfo           { return &Bool{} }
func (t *Bool) Type() slang.TokenType { return slang.TokenBOOL }
func (t *Bool) Value() string         { return StrBOOL }
