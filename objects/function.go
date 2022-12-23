package objects

import (
	"fmt"
	"strings"

	"github.com/jamestunnell/slang"
)

type Function struct {
	Params []string
	Body   slang.Statement
	Env    *slang.Environment
}

const ClassFUNCTION = "Function"

var funClass = NewBuiltInClass(ClassFUNCTION)

func NewFunction(
	params []string, body slang.Statement, env *slang.Environment) *Function {
	return &Function{
		Params: params,
		Body:   body,
		Env:    env,
	}
}

func (obj *Function) Class() slang.Class {
	return funClass
}

func (obj *Function) Inspect() string {
	paramsStr := strings.Join(obj.Params, ", ")

	return fmt.Sprintf("func(%s){...}", paramsStr)
}

func (obj *Function) Truthy() bool {
	return true
}

func (obj *Function) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	// an added instance method would override a standard one
	if m, found := funClass.GetInstanceMethod(methodName); found {
		return m.Run(args)
	}

	switch methodName {
	case slang.MethodCALL:
		return obj.call(args)
	}

	err := slang.NewErrMethodUndefined(methodName, ClassFUNCTION)

	return NULL(), err
}

func (obj *Function) call(args []slang.Object) (slang.Object, error) {
	if len(args) != len(obj.Params) {
		err := slang.NewErrArgCount(len(obj.Params), len(args))

		return NULL(), err
	}

	newEnv := slang.NewEnvironment(obj.Env)

	for i, param := range obj.Params {
		newEnv.Set(param, args[i])
	}

	return obj.Body.Eval(newEnv)
}
