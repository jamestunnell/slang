package slang

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, obj Object)
}

type EnvBase struct {
	parent Environment

	objects map[string]Object
}

func NewEnvBase(parent Environment) *EnvBase {
	return &EnvBase{
		parent:  parent,
		objects: map[string]Object{},
	}
}

func (env *EnvBase) Get(name string) (Object, bool) {
	if obj, found := env.objects[name]; found {
		return obj, true
	}

	if env.parent != nil {
		return env.parent.Get(name)
	}

	return nil, false
}

func (env *EnvBase) Set(name string, obj Object) {
	env.objects[name] = obj
}
