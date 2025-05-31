package slang

import "encoding/json"

type TypeType int

const (
	TypeARRAY TypeType = iota
	TypeBOOLEAN
	TypeEMPTY
	TypeERROR
	TypeFLOAT
	TypeINTEGER
	TypeMAP
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
	StrTypeARRAY   = "ARRAY"
	StrTypeBOOLEAN = "BOOLEAN"
	StrTypeEMPTY   = "EMPTY"
	StrTypeERROR   = "ERROR"
	StrTypeFLOAT   = "FLOAT"
	StrTypeINTEGER = "INTEGER"
	StrTypeMAP     = "MAP"
	StrTypeSTRING  = "STRING"
	StrTypeSTRUCT  = "STRUCT"
)

func TypesEqual(a, b Type) bool {
	return a.IsEqual(b)
}

func ParseTypeTypeStr(s string) (TypeType, bool) {
	var tt TypeType

	switch s {
	case StrTypeARRAY:
		tt = TypeARRAY
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
	case StrTypeMAP:
		tt = TypeMAP
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
	case TypeARRAY:
		s = StrTypeARRAY
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
	case TypeMAP:
		s = StrTypeMAP
	case TypeSTRING:
		s = StrTypeSTRING
	case TypeSTRUCT:
		s = StrTypeSTRUCT
	}

	return s
}
