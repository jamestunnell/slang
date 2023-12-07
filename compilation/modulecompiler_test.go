package compilation_test

import (
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/compilation"
	"github.com/stretchr/testify/assert"
)

func TestModuleCompiler_FirstPass_HappyPath(t *testing.T) {
	stmts := []slang.Statement{
		statements.NewClass("MyClass", "",
			statements.NewMethod("MyMethod", ast.NewFunction(
				[]*ast.Param{}, []string{"bool"}),
			),
			statements.NewField("MyField", "string"),
		),
		statements.NewFunc("MyFunc", ast.NewFunction(
			[]*ast.Param{}, []string{"bool"})),
	}
	expectedSyms := []compilation.Symbol{
		compilation.NewSymbol("", "MyClass", compilation.SymbolCLASS),
		compilation.NewSymbol("MyClass", "MyMethod", compilation.SymbolMETHOD),
		compilation.NewSymbol("MyClass", "MyField", compilation.SymbolFIELD),
		compilation.NewSymbol("", "MyFunc", compilation.SymbolFUNC),
	}
	stmtSeq := compilation.NewStmtSeq(stmts)
	mc := compilation.NewModuleCompiler(stmtSeq)

	err := mc.FirstPass()

	assert.NoError(t, err)

	symbols := mc.CollectSymbols()

	assert.ElementsMatch(t, expectedSyms, symbols)
}
