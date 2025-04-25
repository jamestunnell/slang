package models

import (
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/objects"
)

type EvaluateExprArgs struct {
	Expr *expressions.Expression
}

type EvaluateExprReply struct {
	Error  string
	Object *objects.Object
}
