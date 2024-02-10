package tokens

import "github.com/jamestunnell/slang"

type Continue struct{}

const StrCONTINUE = "continue"

func CONTINUE() slang.TokenInfo           { return &Continue{} }
func (t *Continue) Type() slang.TokenType { return slang.TokenCONTINUE }
func (t *Continue) Value() string         { return StrCONTINUE }
