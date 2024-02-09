package slang

import "fmt"

type Param interface {
	GetName() string
	GetType() Type
}

func ParamString(p Param) string {
	return fmt.Sprintf("%s %s", p.GetName(), p.GetType())
}
