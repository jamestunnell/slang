package slang

import (
	"encoding/json"
)

type StatementType int

type Statement interface {
	Render(level int, w CodeWriter)
	GetName() (string, bool)
	GetType() StatementType
	GetComment() string
	IsEqual(Statement) bool
	// Eval(env *objecsts.Environment) (objects.Object, error)
}

const (
	StatementASSIGN StatementType = iota
	StatementBREAK
	StatementCOMMENT
	StatementCONST
	StatementCONTINUE
	StatementEXPRESSION
	StatementFIELD
	StatementFOREACH
	StatementINTERFACE
	StatementVAR
	StatementFUNC
	StatementIF
	StatementUSE
	StatementIFELSE
	StatementRETURN
	StatementSTRUCT

	StrStatementASSIGN     = "ASSIGN"
	StrStatementBREAK      = "BREAK"
	StrStatementCOMMENT    = "COMMENT"
	StrStatementCONST      = "CONST"
	StrStatementCONTINUE   = "CONTINUE"
	StrStatementEXPRESSION = "EXPRESSION"
	StrStatementFIELD      = "FIELD"
	StrStatementFOREACH    = "FOREACH"
	StrStatementFUNC       = "FUNC"
	StrStatementINTERFACE  = "INTERFACE"
	StrStatementIF         = "IF"
	StrStatementIFELSE     = "IFELSE"
	StrStatementUSE        = "USE"
	StrStatementVAR        = "VAR"
	StrStatementRETURN     = "RETURN"
	StrStatementSTRUCT     = "STRUCT"
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
	case StrStatementFIELD:
		st = StatementFIELD
	case StrStatementFOREACH:
		st = StatementFOREACH
	case StrStatementFUNC:
		st = StatementFUNC
	case StrStatementINTERFACE:
		st = StatementINTERFACE
	case StrStatementIF:
		st = StatementIF
	case StrStatementIFELSE:
		st = StatementIFELSE
	case StrStatementUSE:
		st = StatementUSE
	case StrStatementRETURN:
		st = StatementRETURN
	case StrStatementSTRUCT:
		st = StatementSTRUCT
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
	case StatementFIELD:
		str = StrStatementFIELD
	case StatementFOREACH:
		str = StrStatementFOREACH
	case StatementFUNC:
		str = StrStatementFUNC
	case StatementINTERFACE:
		str = StrStatementINTERFACE
	case StatementIF:
		str = StrStatementIF
	case StatementIFELSE:
		str = StrStatementIFELSE
	case StatementUSE:
		str = StrStatementUSE
	case StatementRETURN:
		str = StrStatementRETURN
	case StatementSTRUCT:
		str = StrStatementSTRUCT
	case StatementVAR:
		str = StrStatementVAR
	}

	return str
}
