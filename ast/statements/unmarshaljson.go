package statements

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/tidwall/gjson"
)

var errTypeNotFound = errors.New("type not found")
var errTypeNotString = errors.New("type not a string")

func UnmarshalJSON(d []byte) (slang.Statement, error) {
	result := gjson.GetBytes(d, "type")
	if !result.Exists() {
		return nil, errTypeNotFound
	}

	if result.Type != gjson.String {
		return nil, errTypeNotString
	}

	var st slang.Statement

	typeStr := result.String()
	switch typeStr {
	case slang.StrASSIGN:
		var ass Assign
		if err := json.Unmarshal(d, &ass); err != nil {
			return nil, fmt.Errorf("failed to unmarshal as assign statement: %w", err)
		}

		st = &ass
	case slang.StrEXPRESSION:
		var expr Expression
		if err := json.Unmarshal(d, &expr); err != nil {
			return nil, fmt.Errorf("failed to unmarshal as expression statement: %w", err)
		}

		st = &expr
	case slang.StrFUNC:
		var fn Func
		if err := json.Unmarshal(d, &fn); err != nil {
			return nil, fmt.Errorf("failed to unmarshal as func statement: %w", err)
		}

		st = &fn
	case slang.StrRETURN:
		var ret Return
		if err := json.Unmarshal(d, &ret); err != nil {
			return nil, fmt.Errorf("failed to unmarshal as return statement: %w", err)
		}

		st = &ret
	case slang.StrUSE:
		var use Use
		if err := json.Unmarshal(d, &use); err != nil {
			return nil, fmt.Errorf("failed to unmarshal as use statement: %w", err)
		}
		st = &use
	}

	if st == nil {
		return nil, customerrs.NewErrUnknownType(typeStr)
	}

	return st, nil
}
