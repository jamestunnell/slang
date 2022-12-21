package slang

type Environment struct {
	Parent *Environment

	objects map[string]Object
}

func NewEnvironment(Parent *Environment) *Environment {
	return &Environment{
		Parent:  Parent,
		objects: map[string]Object{},
	}
}

func (env *Environment) Get(name string) (Object, bool) {
	if obj, found := env.objects[name]; found {
		return obj, true
	}

	if env.Parent != nil {
		return env.Parent.Get(name)
	}

	return nil, false
}

func (env *Environment) Set(name string, obj Object) {
	env.objects[name] = obj
}
