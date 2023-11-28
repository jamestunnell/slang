package tokens

import "github.com/jamestunnell/slang"

type DollarLBrace struct{}

func DOLLARLBRACE() slang.TokenInfo           { return &DollarLBrace{} }
func (t *DollarLBrace) Type() slang.TokenType { return slang.TokenDOLLARLBRACE }
func (t *DollarLBrace) Value() string         { return "${" }
