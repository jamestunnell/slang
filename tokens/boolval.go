package tokens

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type BoolVal struct{ val bool }

const (
	StrFALSE = "false"
	StrTRUE  = "true"
)

func BOOLVAL(val bool) slang.TokenInfo   { return &BoolVal{val: val} }
func (t *BoolVal) Type() slang.TokenType { return slang.TokenBOOLVAL }
func (t *BoolVal) Value() string         { return strconv.FormatBool(t.val) }
