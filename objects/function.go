package objects

import (
	"fmt"
	"strings"

	"github.com/akrennmair/slice"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Function struct {
	Params []slang.Param
	Body   slang.Statement
	Env    *Environment
}

const ClassFUNCTION = "Function"

var funClass = NewBuiltInClass(ClassFUNCTION)

func NewFunction(
	params []slang.Param, body slang.Statement, env *Environment) *Function {
	return &Function{
		Params: params,
		Body:   body,
		Env:    env,
	}
}

func (obj *Function) Class() Class {
	return funClass
}

func (obj *Function) Inspect() string {
	paramStrings := slice.Map(obj.Params, slang.ParamString)
	paramsStr := strings.Join(paramStrings, ", ")

	return fmt.Sprintf("func(%s){...}", paramsStr)
}

func (obj *Function) Truthy() bool {
	return true
}

func (obj *Function) Send(methodName string, args ...Object) (Object, error) {
	// an added instance method would override a standard one
	if m, found := funClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	switch methodName {
	case slang.MethodCALL:
		return obj.call(args)
	}

	err := customerrs.NewErrMethodUndefined(methodName, ClassFUNCTION)

	return NULL(), err
}

func (obj *Function) call(args []Object) (Object, error) {
	if len(args) != len(obj.Params) {
		err := customerrs.NewErrArgCount(len(obj.Params), len(args))

		return NULL(), err
	}

	newEnv := NewEnvironment(obj.Env)

	for i, param := range obj.Params {
		newEnv.Set(param.GetName(), args[i])
	}

	// return obj.Body.Eval(newEnv)

	return NULL(), customerrs.NewErrNotImplemented("Function")
}
