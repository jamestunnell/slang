package slang

type (
	Object interface {
		Inspect() string
		Truthy() bool
		Type() ObjectType
		Send(method string, args ...Object) (Object, error)
	}

	ObjectType int
)

const (
	ObjectBOOL ObjectType = iota
	ObjectFLOAT
	ObjectFUNCTION
	ObjectINTEGER
	ObjectNULL
	ObjectSTRING
)

func (ot ObjectType) String() string {
	var str string

	switch ot {
	case ObjectBOOL:
		str = "BOOL"
	case ObjectFLOAT:
		str = "FLOAT"
	case ObjectFUNCTION:
		str = "FUNCTION"
	case ObjectINTEGER:
		str = "INTEGER"
	case ObjectNULL:
		str = "NULL"
	case ObjectSTRING:
		str = "STRING"
	}
	return str
}
