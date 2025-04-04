package expressions

import (
	"encoding/json"
	"fmt"

	"github.com/jamestunnell/slang"
)

type Base struct {
	ExprType slang.ExprType
}

type baseJSON struct {
	TypeStr string `json:"type"`
}

func NewBase(typ slang.ExprType) *Base {
	return &Base{ExprType: typ}
}

func (b *Base) GetType() slang.ExprType {
	return b.ExprType
}
func (b *Base) MarshalJSON() ([]byte, error) {
	j := &baseJSON{
		TypeStr: b.ExprType.String(),
	}

	return json.Marshal(j)
}

func (b *Base) UnmarshalJSON(d []byte) error {
	var j baseJSON

	if err := json.Unmarshal(d, &j); err != nil {
		return err
	}

	et, ok := slang.ParseExprTypeStr(j.TypeStr)
	if !ok {
		return fmt.Errorf("unknown statement type string '%s'", j.TypeStr)
	}

	b.ExprType = et

	return nil
}
