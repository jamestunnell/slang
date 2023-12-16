package compilation_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/compilation"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestModuleCompiler_FirstPass_HappyPath(t *testing.T) {
	input := `
	class MyClass {
		method MyMethod() bool {
			return true
		}

		field MyField string
	}

	func MyFunc() bool {
		return false
	}
	`

	l := lexer.New(bufio.NewReader(strings.NewReader(input)))
	toks := parsing.NewTokenSeq(l)
	p := parsing.NewFileParser()

	require.True(t, p.Run(toks))

	moduleSymbol := slang.NewRootSymbol("abc", slang.SymbolMODULE)
	expected := map[string]slang.SymbolType{
		"abc.MyClass":          slang.SymbolCLASS,
		"abc.MyClass.MyMethod": slang.SymbolMETHOD,
		"abc.MyClass.MyField":  slang.SymbolFIELD,
		"abc.MyFunc":           slang.SymbolFUNC,
	}
	stmtSeq := compilation.NewStmtSeq(p.Statements)
	mc := compilation.NewModuleCompiler(moduleSymbol, stmtSeq)

	err := mc.FirstPass()

	assert.NoError(t, err)

	childSymbols := mc.ChildSymbols()

	require.Len(t, childSymbols, len(expected))

	for _, sym := range childSymbols {
		scopedName := sym.ScopedName()

		if assert.Contains(t, expected, scopedName) {
			assert.Equal(t, expected[scopedName], sym.Type)
		}
	}
}
