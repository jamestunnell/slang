package objects

import (
	"bytes"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type CompiledFunc struct {
	Instructions []byte
	NumLocals    int
}

const ClassCOMPILEDFUNC = "CompiledFunc"

// var funClass = NewBuiltInClass(ClassFUNCTION)

func NewCompiledFunc(instructions []byte, numLocals int) *CompiledFunc {
	return &CompiledFunc{
		Instructions: instructions,
		NumLocals:    numLocals,
	}
}

func (obj *CompiledFunc) Equal(other slang.Object) bool {
	obj2, ok := other.(*CompiledFunc)
	if !ok {
		return false
	}

	if obj.NumLocals != obj2.NumLocals {
		return false
	}

	return bytes.Equal(obj.Instructions, obj2.Instructions)
}

func (obj *CompiledFunc) Inspect() string {
	return "func(...){...}"
}

func (obj *CompiledFunc) Send(methodName string, args ...slang.Object) (slang.Object, error) {
	switch methodName {
	case slang.MethodCALL:
		return obj.call(args)
	}

	err := customerrs.NewErrMethodUndefined(methodName, ClassCOMPILEDFUNC)

	return NULL(), err
}

func (obj *CompiledFunc) call(args []slang.Object) (slang.Object, error) {
	// if len(args) != len(obj.Params) {
	// 	err := customerrs.NewErrArgCount(len(obj.Params), len(args))

	// 	return NULL(), err
	// }

	// newEnv := NewEnvironment(obj.Env)

	// for i, param := range obj.Params {
	// 	newEnv.Set(param.GetName(), args[i])
	// }

	// // return obj.Body.Eval(newEnv)

	return NULL(), customerrs.NewErrNotImplemented("CompiledFunc")
}
