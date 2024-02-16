package slang

import "encoding/json"

type StatementType int

type Statement interface {
	GetType() StatementType
	IsEqual(Statement) bool
	Eval(env Environment) (Objects, error)
}

const (
	StatementASSIGN StatementType = iota
	StatementBLOCK
	StatementBREAK
	StatementCLASS
	StatementCONST
	StatementCONTINUE
	StatementEXPRESSION
	StatementFIELD
	StatementFOREACH
	StatementVAR
	StatementFUNC
	StatementIF
	StatementIFELSE
	StatementMETHOD
	StatementRETURN
	StatementRETURNVAL
	StatementUSE

	StrASSIGN     = "ASSIGN"
	StrCLASS      = "CLASS"
	StrCONST      = "CONST"
	StrEXPRESSION = "EXPRESSION"
	StrIF         = "IF"
	StrIFELSE     = "IFELSE"
	StrFIELD      = "FIELD"
	StrFUNC       = "FUNC"
	StrVAR        = "VAR"
	StrMETHOD     = "METHOD"
	StrRETURN     = "RETURN"
	StrRETURNVAL  = "RETURNVAL"
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
	case StatementRETURNVAL:
		str = StrRETURNVAL
	case StatementUSE:
		str = StrUSE
	case StatementVAR:
		str = StrVAR
	}

	return str
}
