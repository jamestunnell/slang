package objects

import (
	"github.com/jamestunnell/slang"
)

type Base struct {
	*slang.Env

	Class slang.Class
}

func NewBase(cls slang.Class) *Base {
	return &Base{
		Env:   slang.NewEnv(cls),
		Class: cls,
	}
}

func (obj *Base) GetClass() slang.Class {
	return obj.Class
}
