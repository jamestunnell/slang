package objects

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/customerrs"
)

func Call(obj slang.Object, method string, args ...slang.Object) ([]slang.Object, error) {
	m, found := obj.GetClass().GetMethod(method)
	if !found {
		err := customerrs.NewErrMethodUndefined(method, obj.GetClass().GetType().String())

		return []slang.Object{}, err
	}

	// TODO: check args vs method param types

	return m(obj, args)
}
