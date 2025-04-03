package tokens

import "github.com/jamestunnell/slang"

type Err struct{}

const StrERR = "err"

func ERR() slang.TokenInfo           { return &Err{} }
func (t *Err) Type() slang.TokenType { return slang.TokenERR }
func (t *Err) Value() string         { return StrERR }
