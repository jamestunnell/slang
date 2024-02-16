package slang

type Environment interface {
	Get(string) (Object, bool)
	Set(string, Object)
}

type Env struct {
	parent Environment

	objects map[string]Object
}

func NewEnv(parent Environment) *Env {
	return &Env{
		parent:  parent,
		objects: map[string]Object{},
	}
}

func (env *Env) Get(name string) (Object, bool) {
	if obj, found := env.objects[name]; found {
		return obj, true
	}

	if env.parent != nil {
		return env.parent.Get(name)
	}

	return nil, false
}

func (env *Env) Set(name string, obj Object) {
	env.objects[name] = obj
}
