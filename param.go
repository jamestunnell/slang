package slang

import "fmt"

type Param interface {
	GetName() string
	GetType() string
}

func ParamString(p Param) string {
	return fmt.Sprintf("%s %s", p.GetName(), p.GetType())
}
