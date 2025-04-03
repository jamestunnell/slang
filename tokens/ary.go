package tokens

import "github.com/jamestunnell/slang"

type Ary struct{}

const StrARY = "ary"

func ARY() slang.TokenInfo           { return &Ary{} }
func (t *Ary) Type() slang.TokenType { return slang.TokenARY }
func (t *Ary) Value() string         { return StrARY }
