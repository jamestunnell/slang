package compiler

import (
	"errors"
	"fmt"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/runtime"
	"github.com/jamestunnell/slang/runtime/instructions"
	"github.com/jamestunnell/slang/runtime/objects"
)

type Compiler struct {
	instructions    runtime.Instructions
	globalCollector *GlobalCollector
	symbolTable     *SymbolTable
}

type GlobalCollector struct {
	Constants  []slang.Object
	MaxGlobals int
}

var errTooManyConstants = errors.New("too many constants")

func New() *Compiler {
	return &Compiler{
		instructions: runtime.Instructions{},
		symbolTable:  NewSymbolTable(),
		globalCollector: &GlobalCollector{
			Constants:  []slang.Object{},
			MaxGlobals: 0,
		},
	}
}

func NewChild(parent *Compiler) *Compiler {
	return &Compiler{
		instructions:    runtime.Instructions{},
		globalCollector: parent.globalCollector,
		symbolTable:     NewEnclosedSymbolTable(parent.symbolTable),
	}
}

func (c *Compiler) Instructions() runtime.Instructions {
	return c.instructions
}

func (c *Compiler) MakeBytecode() *runtime.Bytecode {
	return &runtime.Bytecode{
		Instructions: c.instructions.Assemble(),
		Constants:    c.globalCollector.Constants,
		MaxGlobals:   c.globalCollector.MaxGlobals,
	}
}

func (c *Compiler) Symbols() *SymbolTable {
	return c.symbolTable
}

func (c *Compiler) ProcessStmt(stmt slang.Statement) error {
	var err error

	switch stmt.Type() {
	case slang.StatementASSIGN:
		err = c.proccessAssignStmt(stmt.(*statements.Assign))
	case slang.StatementRETURN:
		err = c.proccessReturnStmt(stmt.(*statements.Return))
	case slang.StatementRETURNVAL:
		err = c.proccessReturnValStmt(stmt.(*statements.ReturnVal))
	case slang.StatementBLOCK:
		for _, stmt := range stmt.(*statements.Block).Statements {
			if err = c.ProcessStmt(stmt); err != nil {
				break
			}
		}
	case slang.StatementFUNC:
		err = c.processFuncStmt(stmt.(*statements.Func))
	case slang.StatementIF:
		err = c.processIfStmt(stmt.(*statements.If))
	case slang.StatementIFELSE:
		err = c.processIfElseStmt(stmt.(*statements.IfElse))
	case slang.StatementEXPRESSION:
		if err = c.processExpr(stmt.(*statements.Expression).Value); err != nil {
			break
		}

		c.addInstr(instructions.NewPop())
	default:
		err = fmt.Errorf("statement type %s not supported", stmt.Type().String())
	}

	return err
}

func (c *Compiler) addInstr(instr *runtime.Instruction) {
	c.instructions = append(c.instructions, instr)
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

	c.storeSymbol(sym)

	return nil
}

func (c *Compiler) proccessReturnStmt(stmt *statements.Return) error {
	c.addInstr(instructions.NewReturn())

	return nil
}

func (c *Compiler) proccessReturnValStmt(stmt *statements.ReturnVal) error {
	if err := c.processExpr(stmt.Value); err != nil {
		return fmt.Errorf("failed to process assign value expr: %w", err)
	}

	c.addInstr(instructions.NewReturnVal())

	return nil
}

func (c *Compiler) processFuncStmt(funcStmt *statements.Func) error {
	compiledFn, numFree, err := c.processFunction(funcStmt.Function, funcStmt.Name)
	if err != nil {
		return fmt.Errorf("failed to process function: %w", err)
	}

	idx, ok := c.addConst(compiledFn)
	if !ok {
		return errTooManyConstants
	}

	c.addInstr(instructions.NewClosure(idx, numFree))

	c.storeSymbol(c.symbolTable.Define(funcStmt.Name))

	return nil
}

func (c *Compiler) storeSymbol(s Symbol) {
	switch s.Scope {
	case GlobalScope:
		c.addInstr(instructions.NewSetGlobal(uint16(s.Index)))

		c.globalCollector.MaxGlobals++
	case LocalScope:
		c.addInstr(instructions.NewSetLocal(uint8(s.Index)))
	}
}

func (c *Compiler) loadSymbol(s Symbol) {
	switch s.Scope {
	case GlobalScope:
		c.addInstr(instructions.NewGetGlobal(uint16(s.Index)))
	case LocalScope:
		c.addInstr(instructions.NewGetLocal(uint8(s.Index)))
	case FreeScope:
		c.addInstr(instructions.NewGetFree(uint8(s.Index)))
	case FunctionScope:
		c.addInstr(instructions.NewCurrentClosure())
	}
}

func (c *Compiler) processIfStmt(stmt *statements.If) error {
	if err := c.processExpr(stmt.Condition); err != nil {
		return fmt.Errorf("failed to process if condition expr: %w", err)
	}

	jumpPastIf, targetAfterIf := instructions.NewJumpIfFalse(runtime.DummyJumpTarget)

	c.addInstr(jumpPastIf)

	if err := c.ProcessStmt(stmt.Block); err != nil {
		return fmt.Errorf("failed to process if block: %w", err)
	}

	targetAfterIf.Value = uint64(c.instructions.LengthBytes())

	return nil
}

func (c *Compiler) processIfElseStmt(stmt *statements.IfElse) error {
	if err := c.processExpr(stmt.Condition); err != nil {
		return fmt.Errorf("failed to process if-condition expr: %w", err)
	}

	jumpPastIf, targetAfterIf := instructions.NewJumpIfFalse(runtime.DummyJumpTarget)

	c.addInstr(jumpPastIf)

	if err := c.ProcessStmt(stmt.IfBlock); err != nil {
		return fmt.Errorf("failed to process if block: %w", err)
	}

	jumpPastElse, targetAfterElse := instructions.NewJump(runtime.DummyJumpTarget)

	c.addInstr(jumpPastElse)

	targetAfterIf.Value = uint64(c.instructions.LengthBytes())

	if err := c.ProcessStmt(stmt.ElseBlock); err != nil {
		return fmt.Errorf("failed to process if block: %w", err)
	}

	targetAfterElse.Value = uint64(c.instructions.LengthBytes())

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
	case *expressions.Call:
		err = c.processCall(ee)
	case *expressions.Func:
		err = c.processFuncLiteral(ee)
	default:
		err = fmt.Errorf("expression type %s not supported", expr.Type().String())
	}

	return err
}

func (c *Compiler) processCall(expr *expressions.Call) error {
	// TODO: make sure the call args match function signature

	if err := c.processExpr(expr.Function); err != nil {
		return fmt.Errorf("failed to process func expression for call: %w", err)
	}

	if len(expr.Args) > 255 {
		return fmt.Errorf("arg count %d > 255", len(expr.Args))
	}

	for i, arg := range expr.Args {
		if arg.Name != "" {
			return customerrs.NewErrNotImplemented("keyword args")
		}

		if err := c.processExpr(arg.Value); err != nil {
			return fmt.Errorf("failed to process arg %d expression for call: %w", i, err)
		}
	}

	c.addInstr(instructions.NewCall(uint8(len(expr.Args))))

	return nil
}

func (c *Compiler) processFuncLiteral(expr *expressions.Func) error {
	compiledFn, numFreeVars, err := c.processFunction(expr.Function, "")
	if err != nil {
		return err
	}

	idx, ok := c.addConst(compiledFn)
	if !ok {
		return errTooManyConstants
	}

	c.addInstr(instructions.NewClosure(idx, numFreeVars))

	return nil
}

func (c *Compiler) addConst(obj slang.Object) (uint16, bool) {
	numConsts := len(c.globalCollector.Constants)

	if numConsts == runtime.MaxVMConstants {
		return 0, false
	}

	c.globalCollector.Constants = append(c.globalCollector.Constants, obj)

	return uint16(numConsts), true
}

func (c *Compiler) processFunction(fn *ast.Function, name string) (*objects.CompiledFunc, uint8, error) {
	funcCompiler := NewChild(c)

	if name != "" {
		funcCompiler.symbolTable.DefineFunctionName(name)
	}

	// define function parameters as local variables
	for _, name := range fn.GetParamNames() {
		funcCompiler.symbolTable.Define(name)
	}

	// check number of return values matches function signature
	switch len(fn.ReturnTypes) {
	case 0:
		for _, stmt := range fn.Statements {
			if stmt.Type() == slang.StatementRETURNVAL {
				return nil, 0, fmt.Errorf("function does not return a value")
			}
		}
	case 1:
		for _, stmt := range fn.Statements {
			if stmt.Type() == slang.StatementRETURN {
				return nil, 0, fmt.Errorf("function should return a value")
			}
		}

		if fn.Statements[len(fn.Statements)-1].Type() != slang.StatementRETURNVAL {
			return nil, 0, fmt.Errorf("missing return value")
		}
	default:
		return nil, 0, customerrs.NewErrNotImplemented("multiple return values")
	}

	for i, stmt := range fn.Statements {
		if err := funcCompiler.ProcessStmt(stmt); err != nil {
			err = fmt.Errorf("failed to process function statement %d: %w", i, err)

			return nil, 0, err
		}
	}

	// allow an implicit return if the function does not return a value
	lastStmt := fn.Statements[len(fn.Statements)-1]
	if len(fn.ReturnTypes) == 0 && lastStmt.Type() != slang.StatementRETURN {
		c.addInstr(instructions.NewReturn())
	}

	freeSymbols := funcCompiler.Symbols().FreeSymbols()
	if len(freeSymbols) > 255 {
		return nil, 0, errors.New("free symbol count > 255")
	}

	for _, sym := range freeSymbols {
		c.loadSymbol(sym)
	}

	fmt.Println("compiled function:")
	offset := uint64(0)
	for _, instr := range funcCompiler.instructions {
		fmt.Printf("0x%016x: %s\n", offset, instr.String())

		offset += uint64(instr.LengthBytes())
	}

	assembled := funcCompiler.Instructions().Assemble()

	numLocals := funcCompiler.Symbols().NumDefs()
	compiledFn := objects.NewCompiledFunc(assembled, numLocals)

	return compiledFn, uint8(len(freeSymbols)), nil
}

func (c *Compiler) processIdent(expr *expressions.Identifier) error {
	sym, found := c.symbolTable.Resolve(expr.Name)
	if !found {
		return fmt.Errorf("symbol %s not defined", expr.Name)
	}

	c.loadSymbol(sym)

	return nil
}

func (c *Compiler) processConst(obj slang.Object) error {
	idx, ok := c.addConst(obj)
	if !ok {
		return errTooManyConstants
	}

	c.addInstr(instructions.NewGetConst(idx))

	return nil
}

func (c *Compiler) processUnaryOp(expr *expressions.UnaryOperation) error {
	if err := c.processExpr(expr.Value); err != nil {
		return fmt.Errorf("failed to process unary expr: %w", err)
	}

	var instr *runtime.Instruction

	switch expr.Type() {
	case slang.ExprNOT:
		instr = instructions.NewNot()
	case slang.ExprNEGATIVE:
		instr = instructions.NewNeg()
	default:
		return fmt.Errorf("unknown unary expression type %s", expr.Type().String())
	}

	c.addInstr(instr)

	return nil
}

func (c *Compiler) processBinaryOp(expr *expressions.BinaryOperation) error {
	if err := c.processExpr(expr.Left); err != nil {
		return fmt.Errorf("failed to process binary left-expr: %w", err)
	}

	if err := c.processExpr(expr.Right); err != nil {
		return fmt.Errorf("failed to process binary right-expr: %w", err)
	}

	var instr *runtime.Instruction

	switch expr.Type() {
	case slang.ExprADD:
		instr = instructions.NewAdd()
	case slang.ExprSUBTRACT:
		instr = instructions.NewSub()
	case slang.ExprMULTIPLY:
		instr = instructions.NewMul()
	case slang.ExprDIVIDE:
		instr = instructions.NewDiv()
	case slang.ExprEQUAL:
		instr = instructions.NewEqual()
	case slang.ExprNOTEQUAL:
		instr = instructions.NewNotEqual()
	case slang.ExprLESS:
		instr = instructions.NewLess()
	case slang.ExprLESSEQUAL:
		instr = instructions.NewLessEqual()
	case slang.ExprGREATER:
		instr = instructions.NewGreater()
	case slang.ExprGREATEREQUAL:
		instr = instructions.NewGreaterEqual()
	case slang.ExprAND:
		instr = instructions.NewAnd()
	case slang.ExprOR:
		instr = instructions.NewOr()
	default:
		return fmt.Errorf("unknown binary expression type %s", expr.Type().String())
	}

	c.addInstr(instr)

	return nil
}
