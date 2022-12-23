package objects

import "github.com/jamestunnell/slang"

type BuiltInClass struct {
	name            string
	instanceMethods map[string]slang.Method
}

func NewBuiltInClass(name string) *BuiltInClass {
	return &BuiltInClass{
		name:            name,
		instanceMethods: map[string]slang.Method{},
	}
}

func (c *BuiltInClass) Name() string {
	return c.name
}

func (c *BuiltInClass) GetInstanceMethod(methodName string) (slang.Method, bool) {
	if method, found := c.instanceMethods[methodName]; found {
		return method, true
	}

	return nil, false
}

func (c *BuiltInClass) AddInstanceMethod(methodName string, method slang.Method) {
	c.instanceMethods[methodName] = method
}
