package tokens

import "github.com/jamestunnell/slang"

type Int struct{}

const StrINT = "int"

func INT() slang.TokenInfo           { return &Int{} }
func (t *Int) Type() slang.TokenType { return slang.TokenINT }
func (t *Int) Value() string         { return StrINT }
