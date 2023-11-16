package tokens

import "github.com/jamestunnell/slang"

type Newline struct{}

func NEWLINE() slang.TokenInfo           { return &Newline{} }
func (l *Newline) Type() slang.TokenType { return slang.TokenNEWLINE }
func (l *Newline) Value() string         { return "\n" }
