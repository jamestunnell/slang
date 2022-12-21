package objects

import (
	"fmt"
	"strings"

	"github.com/akrennmair/slice"
	"github.com/jamestunnell/slang"
)

type Function struct {
	Params  []string
	Body    slang.Statement
	Env     *slang.Environment
	methods map[string]*slang.Method
}

func NewFunction(
	params []string, body slang.Statement, env *slang.Environment) *Function {
	f := &Function{
		Params:  params,
		Body:    body,
		Env:     env,
		methods: map[string]*slang.Method{},
	}

	f.methods[slang.MethodCALL] = slang.NewMethod(f.call, params...)
	f.methods[slang.MethodPARAMS] = slang.NewMethod(f.params)

	return f
}

func (obj *Function) Inspect() string {
	paramsStr := strings.Join(obj.Params, ", ")

	return fmt.Sprintf("func(%s){...}", paramsStr)
}

func (obj *Function) Truthy() bool {
	return true
}

func (obj *Function) Type() slang.ObjectType {
	return slang.ObjectFUNCTION
}

func (obj *Function) Methods() map[string]*slang.Method {
	return obj.methods
}

func (obj *Function) call(env *slang.Environment) (slang.Object, error) {
	return obj.Body.Eval(env)
}

func (obj *Function) params(env *slang.Environment) (slang.Object, error) {
	objs := slice.Map(obj.Params, stringToObject)

	return NewArray(objs...), nil
}

func stringToObject(str string) slang.Object {
	return NewString(str)
}
