package tokens

import "github.com/jamestunnell/slang"

type Star struct{}

const (
	StrSTAR = "*"
)

func STAR() slang.TokenInfo           { return &Star{} }
func (t *Star) Type() slang.TokenType { return slang.TokenSTAR }
func (t *Star) Value() string         { return StrSTAR }
