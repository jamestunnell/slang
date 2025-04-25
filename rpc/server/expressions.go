package server

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/objects"
	"github.com/jamestunnell/slang/rpc/models"
)

type Expressions struct {
	VM slang.VirtualMachine
}

func NewExpressions(vm slang.VirtualMachine) *Expressions {
	return &Expressions{VM: vm}
}

func (api *Expressions) Evaluate(args *models.EvaluateExprArgs, reply *models.EvaluateExprReply) error {
	o, err := api.VM.EvaluateExpr(args.Expr)
	if err != nil {
		reply.Object = nil
		reply.Error = fmt.Sprintf("failed to evaluate: %v", err)

		return nil
	}

	obj, ok := o.(*objects.Object)
	if !ok {
		reply.Object = nil
		reply.Error = "result object is not a *objects.Object"

		return nil
	}

	reply.Object = obj
	reply.Error = ""

	return nil
}
