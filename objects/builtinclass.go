package objects

type BuiltInClass struct {
	name            string
	instanceMethods map[string]Method
}

func NewBuiltInClass(name string) *BuiltInClass {
	return &BuiltInClass{
		name:            name,
		instanceMethods: map[string]Method{},
	}
}

func (c *BuiltInClass) Name() string {
	return c.name
}

func (c *BuiltInClass) GetInstanceMethod(methodName string) (Method, bool) {
	if method, found := c.instanceMethods[methodName]; found {
		return method, true
	}

	return nil, false
}

func (c *BuiltInClass) AddInstanceMethod(methodName string, method Method) {
	c.instanceMethods[methodName] = method
}
