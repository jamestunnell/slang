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

func TestFileCompiler_FirstPass_HappyPath(t *testing.T) {
	input := `
	var X int

	class MyClass {
		method MyMethod() bool {
			return true
		}

		field MyField string

		var Y int
	}

	func MyFunc() bool {
		return false
	}
	`

	l := lexer.New(bufio.NewReader(strings.NewReader(input)))
	toks := parsing.NewTokenSeq(l)
	p := parsing.NewFileParser()

	if !assert.True(t, p.Run(toks)) {
		for _, parseErr := range p.GetErrors() {
			t.Logf("parse err at %s: %v", parseErr.Token.Location.String(), parseErr.Error)
		}

		t.FailNow()
	}

	moduleSymbol := slang.NewRootSymbol("abc", slang.SymbolMODULE)
	expected := map[string]slang.SymbolType{
		"abc.X":                slang.SymbolVAR,
		"abc.MyClass":          slang.SymbolCLASS,
		"abc.MyClass.MyMethod": slang.SymbolMETHOD,
		"abc.MyClass.MyField":  slang.SymbolFIELD,
		"abc.MyClass.Y":        slang.SymbolVAR,
		"abc.MyFunc":           slang.SymbolFUNC,
	}
	stmtSeq := compilation.NewStmtSeq(p.Statements)
	mc := compilation.NewFileCompiler(moduleSymbol, stmtSeq)

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
