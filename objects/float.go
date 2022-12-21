package objects

import (
	"strconv"

	"github.com/jamestunnell/slang"
)

type Float struct {
	Value   float64
	methods map[string]*slang.Method
}

func GetFloat(env *slang.Environment, name string) (*Float, error) {
	obj, found := env.Get(name)
	if !found {
		return nil, slang.NewErrObjectNotFound(name)
	}

	f, ok := obj.(*Float)
	if !ok {
		i, ok := obj.(*Float)
		if !ok {
			return nil, slang.NewErrObjectType(slang.ObjectFLOAT, obj.Type())
		}

		f = NewFloat(float64(i.Value))
	}

	return f, nil
}

func NewFloat(val float64) *Float {
	flt := &Float{
		Value:   val,
		methods: map[string]*slang.Method{},
	}

	flt.methods[slang.MethodNEG] = slang.NewMethod(flt.neg)
	flt.methods[slang.MethodABS] = slang.NewMethod(flt.abs)
	flt.methods[slang.MethodADD] = slang.NewMethod(flt.add, ParamOTHER)
	flt.methods[slang.MethodSUB] = slang.NewMethod(flt.sub, ParamOTHER)
	flt.methods[slang.MethodMUL] = slang.NewMethod(flt.mul, ParamOTHER)
	flt.methods[slang.MethodDIV] = slang.NewMethod(flt.div, ParamOTHER)
	flt.methods[slang.MethodEQ] = slang.NewMethod(flt.eq, ParamOTHER)
	flt.methods[slang.MethodNEQ] = slang.NewMethod(flt.neq, ParamOTHER)
	flt.methods[slang.MethodLT] = slang.NewMethod(flt.lt, ParamOTHER)
	flt.methods[slang.MethodLEQ] = slang.NewMethod(flt.leq, ParamOTHER)
	flt.methods[slang.MethodGT] = slang.NewMethod(flt.gt, ParamOTHER)
	flt.methods[slang.MethodGEQ] = slang.NewMethod(flt.geq, ParamOTHER)

	return flt
}

func (obj *Float) Inspect() string {
	return strconv.FormatFloat(obj.Value, 'g', -1, 64)
}

func (obj *Float) Truthy() bool {
	return true
}

func (obj *Float) Type() slang.ObjectType {
	return slang.ObjectFLOAT
}

func (obj *Float) Methods() map[string]*slang.Method {
	return obj.methods
}

func (obj *Float) abs(env *slang.Environment) (slang.Object, error) {
	if obj.Value > 0 {
		return NewFloat(obj.Value), nil
	}

	return NewFloat(-obj.Value), nil
}

func (obj *Float) neg(env *slang.Environment) (slang.Object, error) {
	return NewFloat(-obj.Value), nil
}

func (obj *Float) add(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewFloat(obj.Value + i.Value), nil
}

func (obj *Float) sub(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewFloat(obj.Value - i.Value), nil
}

func (obj *Float) mul(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewFloat(obj.Value * i.Value), nil
}

func (obj *Float) div(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewFloat(obj.Value / i.Value), nil
}

func (obj *Float) eq(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value == i.Value), nil
}

func (obj *Float) neq(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value != i.Value), nil
}

func (obj *Float) lt(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value < i.Value), nil
}

func (obj *Float) leq(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value <= i.Value), nil
}

func (obj *Float) gt(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value > i.Value), nil
}

func (obj *Float) geq(env *slang.Environment) (slang.Object, error) {
	i, err := GetFloat(env, ParamOTHER)
	if err != nil {
		return NULL(), err
	}

	return NewBool(obj.Value >= i.Value), nil
}
