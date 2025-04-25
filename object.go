package slang

type Object interface {
	GetType() ObjectType
	IsEqual(Object) bool
	Inspect() string
	Send(name string, args ...Object) (Object, error)
}

type ObjectType int

const (
	ObjectARRAY ObjectType = iota
	ObjectBOOL
	ObjectERROR
	ObjectFLOAT
	ObjectFUNCTION
	ObjectINTEGER
	// ObjectNULL
	ObjectSTRING
)

func ObjectsEqual(a, b []Object) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, obj := range a {
		if !obj.IsEqual(b[idx]) {
			return false
		}
	}

	return true
}

func (ot ObjectType) String() string {
	var str string

	switch ot {
	case ObjectARRAY:
		str = "ARRAY"
	case ObjectBOOL:
		str = "BOOL"
	case ObjectFLOAT:
		str = "FLOAT"
	case ObjectERROR:
		str = "ERROR"
	case ObjectFUNCTION:
		str = "FUNCTION"
	case ObjectINTEGER:
		str = "INTEGER"
	// case ObjectNULL:
	// 	str = "NULL"
	case ObjectSTRING:
		str = "STRING"
	}
	return str
}
