package slang

type StatementType int

type Statement interface {
	Type() StatementType
	Equal(Statement) bool
	// Eval(env *objecsts.Environment) (objects.Object, error)
}

const (
	StatementASSIGN StatementType = iota
	StatementBLOCK
	StatementEXPRESSION
	StatementFUNC
	StatementRETURN
	StatementUSE
)

func (st StatementType) String() string {
	var str string

	switch st {
	case StatementASSIGN:
		str = "ASSIGN"
	case StatementBLOCK:
		str = "BLOCK"
	case StatementEXPRESSION:
		str = "EXPRESSION"
	case StatementFUNC:
		str = "FUNC"
	case StatementRETURN:
		str = "RETURN"
	case StatementUSE:
		str = "USE"
	}

	return str
}
