package slang

type StatementType int

type Statement interface {
	Type() StatementType
	Equal(Statement) bool
	// Eval(env *objecsts.Environment) (objects.Object, error)
}

const (
	StatementASSIGN StatementType = iota
	StatementEXPRESSION
	StatementFUNC
	StatementRETURN
	StatementUSE

	StrASSIGN     = "ASSIGN"
	StrEXPRESSION = "EXPRESSION"
	StrFUNC       = "FUNC"
	StrRETURN     = "RETURN"
	StrUSE        = "USE"
)

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = StrASSIGN
	case StatementEXPRESSION:
		str = StrEXPRESSION
	case StatementFUNC:
		str = StrFUNC
	case StatementRETURN:
		str = StrRETURN
	case StatementUSE:
		str = StrUSE
	}

	return str
}
