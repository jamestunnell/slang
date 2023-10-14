package tokens

import "github.com/jamestunnell/slang"

type Colon struct{}

func COLON() slang.TokenInfo           { return &Colon{} }
func (t *Colon) Type() slang.TokenType { return slang.TokenCOLON }
func (t *Colon) Value() string         { return ";" }
