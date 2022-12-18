package objects

import (
	"fmt"
	"strings"

	"github.com/jamestunnell/slang"
)

type Function struct {
	ParamNames []string
}

func (obj *Function) Inspect() string {
	paramsStr := strings.Join(obj.ParamNames, ", ")

	return fmt.Sprintf("func(%s){...}", paramsStr)
}

func (obj *Function) Truthy() bool {
	return true
}

func (obj *Function) Type() slang.ObjectType {
	return slang.ObjectFUNCTION
}

func (obj *Function) Send(method string, args ...slang.Object) (slang.Object, error) {
	// switch method {
	// case slang.MethodCALL:

	// }

	err := slang.NewErrMethodUndefined(method, slang.ObjectFUNCTION)

	return NULL(), err
}
