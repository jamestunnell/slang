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
	code        *runtime.Bytecode
	symbolTable *SymbolTable
}

var errTooManyConstants = errors.New("too many constants")

func New() *Compiler {
	return &Compiler{
		code:        runtime.NewBytecode(),
		symbolTable: NewSymbolTable(),
	}
}

func (c *Compiler) GetCode() *runtime.Bytecode {
	return c.code
}

func (c *Compiler) ProcessStmt(stmt slang.Statement) error {
	var err error

	switch stmt.Type() {
	case slang.StatementASSIGN:
		err = c.proccessAssignStmt(stmt.(*statements.Assign))
	case slang.StatementBLOCK:
		for _, stmt := range stmt.(*statements.Block).Statements {
			if err = c.ProcessStmt(stmt); err != nil {
				break
			}
		}
	case slang.StatementIF:
		err = c.processIfStmt(stmt.(*statements.If))
	case slang.StatementIFELSE:
		err = c.processIfElseStmt(stmt.(*statements.IfElse))
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

func (c *Compiler) proccessAssignStmt(stmt *statements.Assign) error {
	if err := c.processExpr(stmt.Value); err != nil {
		return fmt.Errorf("failed to process assign value expr: %w", err)
	}

	if stmt.Target.Type() != slang.ExprIDENTIFIER {
		return fmt.Errorf("non-simple assign target not supported")
	}

	name := stmt.Target.(*expressions.Identifier).Name

	sym, found := c.symbolTable.Resolve(name)
	if !found {
		sym = c.symbolTable.Define(name)
	}

	if sym.Index >= runtime.MaxGlobals {
		return fmt.Errorf("max globals exceeded")
	}

	c.code.AddSetGlobal(uint16(sym.Index))

	return nil
}

func (c *Compiler) processIfStmt(stmt *statements.If) error {
	if err := c.processExpr(stmt.Condition); err != nil {
		return fmt.Errorf("failed to process if condition expr: %w", err)
	}

	fixup := c.code.AddJumpIfFalse()

	if err := c.ProcessStmt(stmt.Block); err != nil {
		return fmt.Errorf("failed to process if block: %w", err)
	}

	fixup(uint64(len(c.code.Instructions)))

	return nil
}

func (c *Compiler) processIfElseStmt(stmt *statements.IfElse) error {
	if err := c.processExpr(stmt.Condition); err != nil {
		return fmt.Errorf("failed to process if-condition expr: %w", err)
	}

	fixupJumpToElse := c.code.AddJumpIfFalse()

	if err := c.ProcessStmt(stmt.IfBlock); err != nil {
		return fmt.Errorf("failed to process if block: %w", err)
	}

	fixupSkipElse := c.code.AddJump()

	fixupJumpToElse(uint64(len(c.code.Instructions)))

	if err := c.ProcessStmt(stmt.ElseBlock); err != nil {
		return fmt.Errorf("failed to process if block: %w", err)
	}

	fixupSkipElse(uint64(len(c.code.Instructions)))

	return nil
}

func (c *Compiler) processExpr(expr slang.Expression) error {
	var err error

	switch ee := expr.(type) {
	case *expressions.Identifier:
		err = c.processIdent(ee)
	case *expressions.UnaryOperation:
		err = c.processUnaryOp(ee)
	case *expressions.BinaryOperation:
		err = c.processBinaryOp(ee)
	case *expressions.Const[bool]:
		err = c.processConst(objects.NewBool(ee.Value))
	case *expressions.Const[int64]:
		err = c.processConst(objects.NewInt(ee.Value))
	case *expressions.Const[float64]:
		err = c.processConst(objects.NewFloat(ee.Value))
	case *expressions.Const[string]:
		err = c.processConst(objects.NewString(ee.Value))
	default:
		err = fmt.Errorf("expression type %s not supported", expr.Type().String())
	}

	return err
}

func (c *Compiler) processIdent(expr *expressions.Identifier) error {
	sym, found := c.symbolTable.Resolve(expr.Name)
	if !found {
		return fmt.Errorf("symbol %s not defined", expr.Name)
	}

	c.code.AddGetGlobal(uint16(sym.Index))

	return nil
}

func (c *Compiler) processConst(obj slang.Object) error {
	idx, ok := c.code.AddConstant(obj)
	if !ok {
		return errTooManyConstants
	}

	c.code.AddInstructionUint16Operands(runtime.OpCONST, idx)

	return nil
}

func (c *Compiler) processUnaryOp(expr *expressions.UnaryOperation) error {
	if err := c.processExpr(expr.Value); err != nil {
		return fmt.Errorf("failed to process unary expr: %w", err)
	}

	var opcode runtime.Opcode

	switch expr.Type() {
	case slang.ExprNOT:
		opcode = runtime.OpNOT
	case slang.ExprNEGATIVE:
		opcode = runtime.OpNEG
	default:
		return fmt.Errorf("unknown unary expression type %s", expr.Type().String())
	}

	c.code.AddInstructionNoOperands(opcode)

	return nil
}

func (c *Compiler) processBinaryOp(expr *expressions.BinaryOperation) error {
	if err := c.processExpr(expr.Left); err != nil {
		return fmt.Errorf("failed to process binary left-expr: %w", err)
	}

	if err := c.processExpr(expr.Right); err != nil {
		return fmt.Errorf("failed to process binary right-expr: %w", err)
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
	case slang.ExprAND:
		opcode = runtime.OpAND
	case slang.ExprOR:
		opcode = runtime.OpOR
	default:
		return fmt.Errorf("unknown binary expression type %s", expr.Type().String())
	}

	c.code.AddInstructionNoOperands(opcode)

	return nil
}
