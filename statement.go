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
	StatementBREAK
	StatementCLASS
	StatementCONST
	StatementCONTINUE
	StatementEXPRESSION
	StatementCLASSFIELD
	StatementFOREACH
	StatementVAR
	StatementFUNC
	StatementIF
	StatementIFELSE
	StatementMETHOD
	StatementRETURN
	StatementRETURNVAL
	StatementSTRUCT
	StatementSTRUCTFIELD
	StatementUSE

	StrASSIGN      = "ASSIGN"
	StrCLASS       = "CLASS"
	StrCLASSFIELD  = "CLASSFIELD"
	StrCONST       = "CONST"
	StrEXPRESSION  = "EXPRESSION"
	StrIF          = "IF"
	StrIFELSE      = "IFELSE"
	StrFUNC        = "FUNC"
	StrVAR         = "VAR"
	StrMETHOD      = "METHOD"
	StrRETURN      = "RETURN"
	StrRETURNVAL   = "RETURNVAL"
	StrSTRUCT      = "STRUCT"
	StrSTRUCTFIELD = "STRUCTFIELD"
	StrUSE         = "USE"
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
	case StatementCLASSFIELD:
		str = StrCLASSFIELD
	case StatementEXPRESSION:
		str = StrEXPRESSION
	case StatementIF:
		str = StrIF
	case StatementIFELSE:
		str = StrIFELSE
	case StatementFUNC:
		str = StrFUNC
	case StatementMETHOD:
		str = StrMETHOD
	case StatementRETURN:
		str = StrRETURN
	case StatementRETURNVAL:
		str = StrRETURNVAL
	case StatementSTRUCT:
		str = StrSTRUCT
	case StatementSTRUCTFIELD:
		str = StrSTRUCTFIELD
	case StatementUSE:
		str = StrUSE
	case StatementVAR:
		str = StrVAR
	}

	return str
}
