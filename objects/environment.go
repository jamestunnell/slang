package objects

// type Environment struct {
// 	parent *Environment

// 	objects map[string]Object
// }

// func NewEnvironment(parent *Environment) *Environment {
// 	return &Environment{
// 		parent:  parent,
// 		objects: map[string]Object{},
// 	}
// }

// func (env *Environment) Get(name string) (Object, bool) {
// 	if obj, found := env.objects[name]; found {
// 		return obj, true
// 	}

// 	if env.parent != nil {
// 		return env.parent.Get(name)
// 	}

// 	return nil, false
// }

// func (env *Environment) Set(name string, obj Object) {
// 	env.objects[name] = obj
// }
