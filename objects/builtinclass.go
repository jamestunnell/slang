package objects

import (
	"github.com/jamestunnell/slang"
)

type BuiltInClass struct {
	*slang.Env

	Type    slang.Type
	Methods map[string]slang.MethodFunc
}

func NewBuiltInClass(
	typ slang.Type,
	methods map[string]slang.MethodFunc,
) *BuiltInClass {
	return &BuiltInClass{
		Type:    typ,
		Methods: methods,
	}
}

func (c *BuiltInClass) GetType() slang.Type {
	return c.Type
}

func (c *BuiltInClass) GetMethod(name string) (slang.MethodFunc, bool) {
	if m, found := c.Methods[name]; found {
		return m, true
	}

	return nil, false
}
