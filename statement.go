package slang

import "encoding/json"

type StatementType int

type Statement interface {
	Type() StatementType
	Equal(Statement) bool
	// Eval(env *objecsts.Environment) (objects.Object, error)
}

const (
	StatementASSIGN StatementType = iota
	StatementCLASS
	StatementEXPRESSION
	StatementFIELD
	StatementFUNC
	StatementMETHOD
	StatementRETURN
	StatementUSE

	StrASSIGN     = "ASSIGN"
	StrCLASS      = "CLASS"
	StrEXPRESSION = "EXPRESSION"
	StrFIELD      = "FIELD"
	StrFUNC       = "FUNC"
	StrMETHOD     = "METHOD"
	StrRETURN     = "RETURN"
	StrUSE        = "USE"
)

func (st StatementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.String())
}

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = StrASSIGN
	case StatementCLASS:
		str = StrCLASS
	case StatementEXPRESSION:
		str = StrEXPRESSION
	case StatementFIELD:
		str = StrFIELD
	case StatementFUNC:
		str = StrFUNC
	case StatementMETHOD:
		str = StrMETHOD
	case StatementRETURN:
		str = StrRETURN
	case StatementUSE:
		str = StrUSE
	}

	return str
}
