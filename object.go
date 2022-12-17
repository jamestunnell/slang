package slang

type Object interface {
	Inspect() string
	Truthy() bool
	Type() ObjectType
	Send(method string, args ...Object) (Object, error)
}

type ObjectType int

const (
	ObjectBOOL ObjectType = iota
	ObjectFLOAT
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
	case ObjectINTEGER:
		str = "INTEGER"
	case ObjectNULL:
		str = "NULL"
	case ObjectSTRING:
		str = "STRING"
	}
	return str
}
