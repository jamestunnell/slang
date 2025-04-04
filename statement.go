package slang

import "encoding/json"

type StatementType int

type Statement interface {
	// SetComment(lines []string)
	GetType() StatementType
	GetComment() string
	IsEqual(Statement) bool
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
	StatementUSE

	StrStatementASSIGN     = "ASSIGN"
	StrStatementBREAK      = "BREAK"
	StrStatementCOMMENT    = "COMMENT"
	StrStatementCONST      = "CONST"
	StrStatementCONTINUE   = "CONTINUE"
	StrStatementEXPRESSION = "EXPRESSION"
	StrStatementFOREACH    = "FOREACH"
	StrStatementIF         = "IF"
	StrStatementIFELSE     = "IFELSE"
	StrStatementFUNC       = "FUNC"
	StrStatementVAR        = "VAR"
	StrStatementMETHOD     = "METHOD"
	StrStatementRETURN     = "RETURN"
	StrStatementRETURNVAL  = "RETURNVAL"
	StrStatementSTRUCT     = "STRUCT"
	StrStatementUSE        = "USE"
)

func StatementsEqual(a, b Statement) bool {
	return a.IsEqual(b)
}

func ParseStatementTypeStr(s string) (StatementType, bool) {
	var st StatementType

	switch s {
	case StrStatementASSIGN:
		st = StatementASSIGN
	case StrStatementBREAK:
		st = StatementBREAK
	case StrStatementCOMMENT:
		st = StatementCOMMENT
	case StrStatementCONST:
		st = StatementCONST
	case StrStatementCONTINUE:
		st = StatementCONTINUE
	case StrStatementEXPRESSION:
		st = StatementEXPRESSION
	case StrStatementFOREACH:
		st = StatementFOREACH
	case StrStatementIF:
		st = StatementIF
	case StrStatementIFELSE:
		st = StatementIFELSE
	case StrStatementFUNC:
		st = StatementFUNC
	case StrStatementMETHOD:
		st = StatementMETHOD
	case StrStatementRETURN:
		st = StatementRETURN
	case StrStatementRETURNVAL:
		st = StatementRETURNVAL
	case StrStatementSTRUCT:
		st = StatementSTRUCT
	case StrStatementUSE:
		st = StatementUSE
	case StrStatementVAR:
		st = StatementVAR
	default:
		return st, false
	}

	return st, true
}

func (st StatementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(st.String())
}

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = StrStatementASSIGN
	case StatementBREAK:
		str = StrStatementBREAK
	case StatementCOMMENT:
		str = StrStatementCOMMENT
	case StatementCONST:
		str = StrStatementCONST
	case StatementCONTINUE:
		str = StrStatementCONTINUE
	case StatementEXPRESSION:
		str = StrStatementEXPRESSION
	case StatementFOREACH:
		str = StrStatementFOREACH
	case StatementIF:
		str = StrStatementIF
	case StatementIFELSE:
		str = StrStatementIFELSE
	case StatementFUNC:
		str = StrStatementFUNC
	case StatementMETHOD:
		str = StrStatementMETHOD
	case StatementRETURN:
		str = StrStatementRETURN
	case StatementRETURNVAL:
		str = StrStatementRETURNVAL
	case StatementSTRUCT:
		str = StrStatementSTRUCT
	case StatementUSE:
		str = StrStatementUSE
	case StatementVAR:
		str = StrStatementVAR
	}

	return str
}
