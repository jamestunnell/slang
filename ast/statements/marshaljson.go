package statements

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/sjson"

	"github.com/jamestunnell/slang"
)

func MarshalJSON(st slang.Statement) ([]byte, error) {
	d, err := json.Marshal(st)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to marshal statement JSON: %w", err)
	}

	d, err = sjson.SetBytes(d, "type", st.Type().String())
	if err != nil {
		return []byte{}, fmt.Errorf("failed to set type: %w", err)
	}

	return d, nil
}
