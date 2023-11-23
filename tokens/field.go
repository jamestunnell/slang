package tokens

import "github.com/jamestunnell/slang"

type Field struct{}

const StrFIELD = "field"

func FIELD() slang.TokenInfo           { return &Field{} }
func (t *Field) Type() slang.TokenType { return slang.TokenFIELD }
func (t *Field) Value() string         { return StrFIELD }
