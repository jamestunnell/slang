package tokens

import "github.com/jamestunnell/slang"

type Map struct{}

const StrMAP = "map"

func MAP() slang.TokenInfo           { return &Map{} }
func (t *Map) Type() slang.TokenType { return slang.TokenMAP }
func (t *Map) Value() string         { return StrMAP }
