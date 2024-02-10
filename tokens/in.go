package tokens

import "github.com/jamestunnell/slang"

type In struct{}

const (
	StrIN = "in"
)

func IN() slang.TokenInfo           { return &In{} }
func (t *In) Type() slang.TokenType { return slang.TokenIN }
func (t *In) Value() string         { return StrIN }
