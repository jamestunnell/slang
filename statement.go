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
	StatementBLOCK
	StatementCLASS
	StatementEXPRESSION
	StatementFIELD
	StatementFUNC
	StatementIF
	StatementIFELSE
	StatementMETHOD
	StatementRETURN
	StatementUSE

	StrASSIGN     = "ASSIGN"
	StrCLASS      = "CLASS"
	StrEXPRESSION = "EXPRESSION"
	StrIF         = "IF"
	StrIFELSE     = "IFELSE"
	StrFIELD      = "FIELD"
	StrFUNC       = "FUNC"
	StrMETHOD     = "METHOD"
	StrRETURN     = "RETURN"
	StrUSE        = "USE"
)

func StatementsEqual(a, b []Statement) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, stmt := range a {
		if !stmt.Equal(b[idx]) {
			return false
		}
	}

	return true
}

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
	case StatementIF:
		str = StrIF
	case StatementIFELSE:
		str = StrIFELSE
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
