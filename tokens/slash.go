package tokens

import "github.com/jamestunnell/slang"

type Slash struct{}

const (
	StrSLASH = "/"
)

func SLASH() slang.TokenInfo           { return &Slash{} }
func (t *Slash) Type() slang.TokenType { return slang.TokenSLASH }
func (t *Slash) Value() string         { return StrSLASH }
