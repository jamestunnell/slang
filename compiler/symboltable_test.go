package compiler_test

import (
	"testing"

	"github.com/jamestunnell/slang/compiler"
	"github.com/stretchr/testify/assert"
)

func TestSymbolTableDefine(t *testing.T) {
	expected := map[string]compiler.Symbol{
		"a": {
			Name:  "a",
			Scope: compiler.GlobalScope,
			Index: 0,
		},
		"b": {
			Name:  "b",
			Scope: compiler.GlobalScope,
			Index: 1,
		},
	}
	global := compiler.NewSymbolTable()

	a := global.Define("a")

	assert.Equal(t, expected["a"], a)

	b := global.Define("b")

	assert.Equal(t, expected["b"], b)
}

func TestSymbolTable(t *testing.T) {
	global := compiler.NewSymbolTable()

	global.Define("a")
	global.Define("b")

	verifyResolves(t, global, "a", 0)
	verifyResolves(t, global, "b", 1)

	firstLocal := compiler.NewEnclosedSymbolTable(global)
	firstLocal.Define("c")
	firstLocal.Define("d")

	verifyResolves(t, firstLocal, "c", 0)
	verifyResolves(t, firstLocal, "d", 1)

	secondLocal := compiler.NewEnclosedSymbolTable(firstLocal)
	secondLocal.Define("e")
	secondLocal.Define("f")

	verifyResolves(t, secondLocal, "c", 0)
	verifyResolves(t, secondLocal, "d", 1)
}

func verifyResolves(
	t *testing.T,
	st *compiler.SymbolTable,
	name string,
	expectedIdx int,
) {
	sym, found := st.Resolve("a")

	if assert.True(t, found) {
		assert.Equal(t, "a", sym.Name)
	}
}
