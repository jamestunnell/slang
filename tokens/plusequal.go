package tokens

import "github.com/jamestunnell/slang"

type PlusEqual struct{}

func PLUSEQUAL() slang.TokenInfo           { return &PlusEqual{} }
func (t *PlusEqual) Type() slang.TokenType { return slang.TokenPLUSEQUAL }
func (t *PlusEqual) Value() string         { return "+=" }
