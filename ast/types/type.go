package types

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/jsonutil"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type Type struct {
	Type slang.TypeType
	Core Core
}

type Core interface {
	IsEqual(Core) bool
	String() string
}

func NewType(typ slang.TypeType, core Core) *Type {
	return &Type{
		Type: typ,
		Core: core,
	}
}

func (t *Type) String() string {
	return t.Core.String()
}

func (t *Type) GetType() slang.TypeType {
	return t.Type
}

func (t *Type) IsEqual(other slang.Type) bool {
	t2, ok := other.(*Type)
	if !ok {
		return false
	}

	if t.Type != t2.Type {
		return false
	}

	return t.Core.IsEqual(t2.Core)
}

func (t *Type) MarshalJSON() ([]byte, error) {
	d, err := json.Marshal(t.Core)
	if err != nil {
		return []byte{}, err
	}

	d, err = sjson.SetBytes(d, "type", t.Type.String())
	if err != nil {
		return []byte{}, err
	}

	return d, nil
}

var (
	errTypeNotFound  = errors.New("type not found")
	errTypeNotString = errors.New("type not a string")
)

func (s *Type) UnmarshalJSON(d []byte) error {
	result := gjson.GetBytes(d, "type")
	if !result.Exists() {
		return errTypeNotFound
	}

	if result.Type != gjson.String {
		return errTypeNotString
	}

	stmtType, ok := slang.ParseTypeTypeStr(result.String())
	if !ok {
		return fmt.Errorf("invalid type type string '%s'", result.String())
	}

	var err error

	var core Core

	switch stmtType {
	case slang.TypeBOOLEAN:
		core = &Bool{}
	case slang.TypeERROR:
		core = &Err{}
	case slang.TypeFLOAT:
		core = &Flt{}
	case slang.TypeINTEGER:
		core = &Int{}
	case slang.TypeSTRING:
		core = &Str{}
	case slang.TypeSTRUCT:
		core, err = jsonutil.UnmarshalAs[Struct](d)
	}

	if err != nil {
		return fmt.Errorf("failed to unmarshal core: %w", err)
	}

	s.Core = core
	s.Type = stmtType

	return nil
}
