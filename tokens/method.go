package tokens

import "github.com/jamestunnell/slang"

type Method struct{}

const StrMETHOD = "method"

func METHOD() slang.TokenInfo           { return &Method{} }
func (t *Method) Type() slang.TokenType { return slang.TokenMETHOD }
func (t *Method) Value() string         { return StrMETHOD }
