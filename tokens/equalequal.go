package tokens

import "github.com/jamestunnell/slang"

type EqualEqual struct{}

const StrEQUALEQUAL = "=="

func EQUALEQUAL() slang.TokenInfo           { return &EqualEqual{} }
func (t *EqualEqual) Type() slang.TokenType { return slang.TokenEQUALEQUAL }
func (t *EqualEqual) Value() string         { return StrEQUALEQUAL }
