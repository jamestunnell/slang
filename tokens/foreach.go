package tokens

import "github.com/jamestunnell/slang"

type ForEach struct{}

const (
	StrFOREACH = "foreach"
)

func FOREACH() slang.TokenInfo           { return &ForEach{} }
func (t *ForEach) Type() slang.TokenType { return slang.TokenFOREACH }
func (t *ForEach) Value() string         { return StrFOREACH }
