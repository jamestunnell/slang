package slang

import "fmt"

type NameType interface {
	GetName() string
	GetType() Type
}

func NameTypeString(nt NameType) string {
	return fmt.Sprintf("%s %s", nt.GetName(), nt.GetType())
}

func NameTypesEqual(nt1, nt2 NameType) bool {
	return (nt1.GetName() == nt2.GetName()) && TypesEqual(nt1.GetType(), nt2.GetType())
}
