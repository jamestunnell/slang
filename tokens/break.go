package tokens

import "github.com/jamestunnell/slang"

type Break struct{}

const StrBREAK = "break"

func BREAK() slang.TokenInfo           { return &Break{} }
func (t *Break) Type() slang.TokenType { return slang.TokenBREAK }
func (t *Break) Value() string         { return StrBREAK }
