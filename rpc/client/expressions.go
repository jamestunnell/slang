package client

import (
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/rpc/models"
)

func (client *Client) EvaluateExpr(e slang.Expression) (slang.Object, error) {
	const method = "Expressions.Evaluate"

	expr, ok := e.(*expressions.Expression)
	if !ok {
		return nil, fmt.Errorf("expression is not a *expressions.Expression")
	}

	args := &models.EvaluateExprArgs{Expr: expr}

	var reply models.EvaluateExprReply

	err := client.rpcClient.Call(method, args, &reply)
	if err != nil {
		return nil, newErrMethodFailed(method, err)
	}

	if reply.Error != "" {
		return nil, fmt.Errorf("eval failed: %s", reply.Error)
	}

	return reply.Object, nil
}

func newErrMethodFailed(method string, err error) *errMethodFailed {
	return &errMethodFailed{
		method: method,
		err:    err,
	}
}

type errMethodFailed struct {
	method string
	err    error
}

func (err *errMethodFailed) Error() string {
	return fmt.Sprintf("method %s failed: %v", err.method, err.err)
}
