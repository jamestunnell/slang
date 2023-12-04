package objects

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

type Closure struct {
	Func     *CompiledFunc
	FreeVars []slang.Object
}

func NewClosure(fn *CompiledFunc, freeVars ...slang.Object) *Closure {
	return &Closure{
		Func:     fn,
		FreeVars: freeVars,
	}
}

func (obj *Closure) Equal(other slang.Object) bool {
	obj2, ok := other.(*Closure)
	if !ok {
		return false
	}

	if !obj.Func.Equal(obj2.Func) {
		return false
	}

	return slang.ObjectsEqual(obj.FreeVars, obj2.FreeVars)
}

func (obj *Closure) Send(name string, args ...slang.Object) (slang.Object, error) {
	return NULL(), customerrs.NewErrNotImplemented("closure methods")
}
