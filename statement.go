package slang

import "encoding/json"

type StatementType int

type Statement interface {
	SetComment(lines []string)

	Type() StatementType
	Equal(Statement) bool
	// Eval(env *objecsts.Environment) (objects.Object, error)
}

const (
	StatementASSIGN StatementType = iota
	StatementBLOCK
	StatementBREAK
	StatementCOMMENT
	StatementCONST
	StatementCONTINUE
	StatementEXPRESSION
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
	StrCOMMENT     = "COMMENT"
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

func StatementsEqual(a, b Statement) bool {
	return a.Equal(b)
}

func (st StatementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.String())
}

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = StrASSIGN
	case StatementCOMMENT:
		str = StrCOMMENT
	case StatementCONST:
		str = StrCONST
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
