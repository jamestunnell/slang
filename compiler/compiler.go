package compiler

import (
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/runtime"
	"github.com/jamestunnell/slang/runtime/objects"
)

type Compiler struct {
	code *runtime.Bytecode
}

var errTooManyConstants = errors.New("too many constants")

func New() *Compiler {
	return &Compiler{
		code: runtime.NewBytecode(),
	}
}

func (c *Compiler) GetCode() *runtime.Bytecode {
	return c.code
}

func (c *Compiler) ProcessStmt(stmt slang.Statement) error {
	var err error

	switch stmt.Type() {
	case slang.StatementEXPRESSION:
		if err = c.processExpr(stmt.(*statements.Expression).Value); err != nil {
			break
		}

		c.code.AddInstructionNoOperands(runtime.OpPOP)
	default:
		err = fmt.Errorf("statement type %s not supported", stmt.Type().String())
	}

	return err
}

func (c *Compiler) processExpr(expr slang.Expression) error {
	var err error

	switch ee := expr.(type) {
	case *expressions.BinaryOperation:
		err = c.processBinOp(ee)
	case *expressions.Bool:
		idx, ok := c.code.AddConstant(objects.NewBool(ee.Value))
		if !ok {
			err = errTooManyConstants

			break
		}

		c.code.AddInstructionUint16Operands(runtime.OpCONST, idx)
	case *expressions.Integer:
		idx, ok := c.code.AddConstant(objects.NewInt(ee.Value))
		if !ok {
			err = errTooManyConstants

			break
		}

		c.code.AddInstructionUint16Operands(runtime.OpCONST, idx)
	case *expressions.Float:
		idx, ok := c.code.AddConstant(objects.NewFloat(ee.Value))
		if !ok {
			err = errTooManyConstants

			break
		}

		c.code.AddInstructionUint16Operands(runtime.OpCONST, idx)
	default:
		err = fmt.Errorf("expression type %s not supported", expr.Type().String())
	}

	return err
}

func (c *Compiler) processBinOp(expr *expressions.BinaryOperation) error {
	if err := c.processExpr(expr.Left); err != nil {
		return fmt.Errorf("failed to process left-expr: %w", err)
	}

	if err := c.processExpr(expr.Right); err != nil {
		return fmt.Errorf("failed to process right-expr: %w", err)
	}

	var opcode runtime.Opcode

	switch expr.Type() {
	case slang.ExprADD:
		opcode = runtime.OpADD
	case slang.ExprSUBTRACT:
		opcode = runtime.OpSUB
	case slang.ExprMULTIPLY:
		opcode = runtime.OpMUL
	case slang.ExprDIVIDE:
		opcode = runtime.OpDIV
	case slang.ExprEQUAL:
		opcode = runtime.OpEQ
	case slang.ExprNOTEQUAL:
		opcode = runtime.OpNEQ
	case slang.ExprLESS:
		opcode = runtime.OpLT
	case slang.ExprLESSEQUAL:
		opcode = runtime.OpLEQ
	case slang.ExprGREATER:
		opcode = runtime.OpGT
	case slang.ExprGREATEREQUAL:
		opcode = runtime.OpGEQ
	default:
		return fmt.Errorf("unknown binary expression type %s", expr.Type().String())
	}

	c.code.AddInstructionNoOperands(opcode)

	return nil
}
