package compiler_test

import (
	"testing"

	"github.com/jamestunnell/slang/compiler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestSymbolTableResolveGlobal(t *testing.T) {
	global := compiler.NewSymbolTable()

	global.Define("a")
	global.Define("b")

	sym, found := global.Resolve("a")

	require.True(t, found)
	assert.Equal(t, "a", sym.Name)

	sym, found = global.Resolve("b")

	require.True(t, found)
	assert.Equal(t, "b", sym.Name)
}
