package tokens

import "github.com/jamestunnell/slang"

type LessLess struct{}

func LESSLESS() slang.TokenInfo           { return &Less{} }
func (t *LessLess) Type() slang.TokenType { return slang.TokenLESSLESS }
func (t *LessLess) Value() string         { return "<<" }
