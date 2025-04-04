package slang

import "fmt"

type NameType interface {
	GetName() string
	GetType() Type
}

func NameTypeString(nt NameType) string {
	return fmt.Sprintf("%s %s", nt.GetName(), nt.GetType())
}
