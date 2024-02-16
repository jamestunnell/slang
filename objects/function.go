package objects

import (
	"fmt"
	"strings"

	"github.com/akrennmair/slice"
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/types"
)

type Function struct {
	*Base

	Params []slang.Param
	Body   slang.Statement
	Env    slang.Environment
}

const funcClassName = "Function"

var funcClass *BuiltInClass

func init() {
	funcClass = NewBuiltInClass(
		types.NewPrimitiveType(funcClassName),
		map[string]slang.MethodFunc{
			slang.MethodCALL: funcCall,
		},
	)
}

func NewFunction(
	params []slang.Param, body slang.Statement, env slang.Environment) *Function {
	return &Function{
		Base:   NewBase(funcClass),
		Params: params,
		Body:   body,
		Env:    env,
	}
}

func (obj *Function) IsEqual(other slang.Object) bool {
	obj2, ok := other.(*Function)
	if !ok {
		return false
	}

	if !obj.Body.Equal(obj2.Body) {
		return false
	}

	return slices.EqualFunc(obj.Params, obj2.Params, func(a, b slang.Param) bool {
		if a.GetName() != b.GetName() {
			return false
		}

		return a.GetType().IsEqual(b.GetType())
	})
}

func (obj *Function) Inspect() string {
	paramStrings := slice.Map(obj.Params, slang.ParamString)
	paramsStr := strings.Join(paramStrings, ", ")

	return fmt.Sprintf("func(%s){...}", paramsStr)
}

func funcCall(obj slang.Object, args slang.Objects) (slang.Objects, error) {
	return obj.(*Function).call(args)
}

func (obj *Function) call(args []slang.Object) (slang.Objects, error) {
	if err := CheckArgCount(args, len(obj.Params)); err != nil {
		return slang.Objects{}, err
	}

	for i, param := range obj.Params {
		argType := args[i].GetClass().GetType()
		if !argType.IsEqual(param.GetType()) {
			err := customerrs.NewErrArgType(param.GetType().String(), argType.String())

			return slang.Objects{}, err
		}
	}

	newEnv := slang.NewEnv(obj.Env)

	for i, param := range obj.Params {
		newEnv.Set(param.GetName(), args[i])
	}

	return obj.Body.Eval(newEnv)
}
