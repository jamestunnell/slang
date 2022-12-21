package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Bool struct {
	Value   bool
	methods map[string]*slang.Method
}

func GetBool(env *slang.Environment, name string) (*Bool, error) {
	obj, found := env.Get(name)
	if !found {
		return nil, slang.NewErrObjectNotFound(name)
	}

	b, ok := obj.(*Bool)
	if !ok {
		return nil, slang.NewErrObjectType(slang.ObjectBOOL, obj.Type())
	}

	return b, nil
}

func NewBool(val bool) slang.Object {
	b := &Bool{
		Value:   val,
		methods: map[string]*slang.Method{},
	}

	b.methods[slang.MethodNOT] = slang.NewMethod(b.not)
	b.methods[slang.MethodEQ] = slang.NewMethod(b.eq, ParamOTHER)
	b.methods[slang.MethodNEQ] = slang.NewMethod(b.neq, ParamOTHER)

	return b
}

func (obj *Bool) Inspect() string {
	return strconv.FormatBool(obj.Value)
}

func (obj *Bool) Truthy() bool {
	return obj.Value
}

func (obj *Bool) Type() slang.ObjectType {
	return slang.ObjectBOOL
}

func (obj *Bool) Methods() map[string]*slang.Method {
	return obj.methods
}

func (obj *Bool) not(env *slang.Environment) (slang.Object, error) {
	return NewBool(!obj.Value), nil
}

func (obj *Bool) eq(env *slang.Environment) (slang.Object, error) {
	b, err := GetBool(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value == b.Value), nil
}

func (obj *Bool) neq(env *slang.Environment) (slang.Object, error) {
	b, err := GetBool(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value != b.Value), nil
}
