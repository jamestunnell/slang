package slang

type (
	Object interface {
		Class() Class
		Inspect() string
		Truthy() bool
		// Type() ObjectType
		Send(method string, args ...Object) (Object, error)
	}

	ObjectType int
)

const (
	ObjectARRAY ObjectType = iota
	ObjectBOOL
	ObjectFLOAT
	ObjectFUNCTION
	ObjectINTEGER
	ObjectNULL
	ObjectSTRING
)

func (ot ObjectType) String() string {
	var str string

	switch ot {
	case ObjectARRAY:
		str = "ARRAY"
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
