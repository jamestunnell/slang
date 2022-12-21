package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Integer struct {
	Value   int64
	methods map[string]*slang.Method
}

func GetInt(env *slang.Environment, name string) (*Integer, error) {
	obj, found := env.Get(name)
	if !found {
		return nil, slang.NewErrObjectNotFound(name)
	}

	i, ok := obj.(*Integer)
	if !ok {
		return nil, slang.NewErrObjectType(slang.ObjectINTEGER, obj.Type())
	}

	return i, nil
}

func NewInteger(val int64) slang.Object {
	i := &Integer{
		Value:   val,
		methods: map[string]*slang.Method{},
	}

	i.methods[slang.MethodNEG] = slang.NewMethod(i.neg)
	i.methods[slang.MethodABS] = slang.NewMethod(i.abs)
	i.methods[slang.MethodADD] = slang.NewMethod(i.add, ParamOTHER)
	i.methods[slang.MethodSUB] = slang.NewMethod(i.sub, ParamOTHER)
	i.methods[slang.MethodMUL] = slang.NewMethod(i.mul, ParamOTHER)
	i.methods[slang.MethodDIV] = slang.NewMethod(i.div, ParamOTHER)
	i.methods[slang.MethodEQ] = slang.NewMethod(i.eq, ParamOTHER)
	i.methods[slang.MethodNEQ] = slang.NewMethod(i.neq, ParamOTHER)
	i.methods[slang.MethodLT] = slang.NewMethod(i.lt, ParamOTHER)
	i.methods[slang.MethodLEQ] = slang.NewMethod(i.leq, ParamOTHER)
	i.methods[slang.MethodGT] = slang.NewMethod(i.gt, ParamOTHER)
	i.methods[slang.MethodGEQ] = slang.NewMethod(i.geq, ParamOTHER)

	return i
}

func (obj *Integer) Inspect() string {
	return strconv.FormatInt(obj.Value, 10)
}

func (obj *Integer) Truthy() bool {
	return true
}

func (obj *Integer) Type() slang.ObjectType {
	return slang.ObjectINTEGER
}

func (obj *Integer) Methods() map[string]*slang.Method {
	return obj.methods
}

func (obj *Integer) abs(env *slang.Environment) (slang.Object, error) {
	if obj.Value > 0 {
		return NewInteger(obj.Value), nil
	}

	return NewInteger(-obj.Value), nil
}

func (obj *Integer) neg(env *slang.Environment) (slang.Object, error) {
	return NewInteger(-obj.Value), nil
}

func (obj *Integer) add(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewInteger(obj.Value + i.Value), nil
}

func (obj *Integer) sub(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewInteger(obj.Value - i.Value), nil
}

func (obj *Integer) mul(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewInteger(obj.Value * i.Value), nil
}

func (obj *Integer) div(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewInteger(obj.Value / i.Value), nil
}

func (obj *Integer) eq(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value == i.Value), nil
}

func (obj *Integer) neq(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value != i.Value), nil
}

func (obj *Integer) lt(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value < i.Value), nil
}

func (obj *Integer) leq(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value <= i.Value), nil
}

func (obj *Integer) gt(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value > i.Value), nil
}

func (obj *Integer) geq(env *slang.Environment) (slang.Object, error) {
	i, err := GetInt(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value >= i.Value), nil
}
