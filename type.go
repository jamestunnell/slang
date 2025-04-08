package slang

import "encoding/json"

type TypeType int

const (
	TypeBOOLEAN TypeType = iota
	TypeEMPTY
	TypeERROR
	TypeFLOAT
	TypeINTEGER
	TypeSTRING
	TypeSTRUCT
)

type Type interface {
	String() string
	GetType() TypeType

	IsEqual(other Type) bool
	// IsArray() bool
	// IsMap() bool
}

const (
	StrTypeBOOLEAN = "BOOLEAN"
	StrTypeEMPTY   = "EMPTY"
	StrTypeERROR   = "ERROR"
	StrTypeFLOAT   = "FLOAT"
	StrTypeINTEGER = "INTEGER"
	StrTypeSTRING  = "STRING"
	StrTypeSTRUCT  = "STRUCT"
)

func TypesEqual(a, b Type) bool {
	return a.IsEqual(b)
}

func ParseTypeTypeStr(s string) (TypeType, bool) {
	var tt TypeType

	switch s {
	case StrTypeBOOLEAN:
		tt = TypeBOOLEAN
	case StrTypeEMPTY:
		tt = TypeEMPTY
	case StrTypeERROR:
		tt = TypeERROR
	case StrTypeFLOAT:
		tt = TypeFLOAT
	case StrTypeINTEGER:
		tt = TypeINTEGER
	case StrTypeSTRING:
		tt = TypeSTRING
	case StrTypeSTRUCT:
		tt = TypeSTRUCT
	default:
		return tt, false
	}

	return tt, true
}

func (tt TypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(tt.String())
}

func (tt TypeType) String() string {
	var s string

	switch tt {
	case TypeBOOLEAN:
		s = StrTypeBOOLEAN
	case TypeEMPTY:
		s = StrTypeEMPTY
	case TypeERROR:
		s = StrTypeERROR
	case TypeFLOAT:
		s = StrTypeFLOAT
	case TypeINTEGER:
		s = StrTypeINTEGER
	case TypeSTRING:
		s = StrTypeSTRING
	case TypeSTRUCT:
		s = StrTypeSTRUCT
	}

	return s
}
