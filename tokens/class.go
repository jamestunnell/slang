package tokens

import "github.com/jamestunnell/slang"

type Class struct{}

const StrCLASS = "class"

func CLASS() slang.TokenInfo           { return &Class{} }
func (t *Class) Type() slang.TokenType { return slang.TokenCLASS }
func (t *Class) Value() string         { return StrCLASS }
