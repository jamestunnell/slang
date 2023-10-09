package tokens

import "github.com/jamestunnell/slang"

type Comment struct {
	val string
}

func COMMENT(val string) slang.TokenInfo { return &Comment{val: val} }
func (t *Comment) Type() slang.TokenType { return slang.TokenCOMMENT }
func (t *Comment) Value() string         { return t.val }
