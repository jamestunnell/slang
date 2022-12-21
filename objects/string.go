package objects

import (
	"github.com/jamestunnell/slang"
)

type String struct {
	Value   string
	methods map[string]*slang.Method
}

const ParamOTHER = "other"

func GetString(env *slang.Environment, name string) (*String, error) {
	obj, found := env.Get(name)
	if !found {
		return nil, slang.NewErrObjectNotFound(name)
	}

	str, ok := obj.(*String)
	if !ok {
		return nil, slang.NewErrObjectType(slang.ObjectSTRING, obj.Type())
	}

	return str, nil
}

func NewString(val string) slang.Object {
	str := &String{
		Value:   val,
		methods: map[string]*slang.Method{},
	}

	str.methods[slang.MethodSIZE] = slang.NewMethod(str.size)
	str.methods[slang.MethodADD] = slang.NewMethod(str.add, ParamOTHER)
	str.methods[slang.MethodEQ] = slang.NewMethod(str.eq, ParamOTHER)
	str.methods[slang.MethodNEQ] = slang.NewMethod(str.neq, ParamOTHER)
	str.methods[slang.MethodLT] = slang.NewMethod(str.lt, ParamOTHER)
	str.methods[slang.MethodGT] = slang.NewMethod(str.gt, ParamOTHER)
	str.methods[slang.MethodLEQ] = slang.NewMethod(str.leq, ParamOTHER)
	str.methods[slang.MethodGEQ] = slang.NewMethod(str.geq, ParamOTHER)

	return str
}

func (obj *String) Inspect() string {
	return obj.Value
}

func (obj *String) Truthy() bool {
	return true
}

func (obj *String) Type() slang.ObjectType {
	return slang.ObjectSTRING
}

func (obj *String) Methods() map[string]*slang.Method {
	return obj.methods
}

func (obj *String) size(env *slang.Environment) (slang.Object, error) {
	sz := NewInteger(int64(len(obj.Value)))
	return sz, nil
}

func (obj *String) add(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewString(obj.Value + str.Value), nil
}

func (obj *String) eq(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value == str.Value), nil
}

func (obj *String) neq(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value != str.Value), nil
}

func (obj *String) lt(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value < str.Value), nil
}

func (obj *String) gt(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value > str.Value), nil
}

func (obj *String) leq(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value <= str.Value), nil
}

func (obj *String) geq(env *slang.Environment) (slang.Object, error) {
	str, err := GetString(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value >= str.Value), nil
}
