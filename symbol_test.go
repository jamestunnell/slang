package slang_test

import (
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/stretchr/testify/assert"
)

func TestSymbol(t *testing.T) {
	s1 := slang.NewRootSymbol("ABC", slang.SymbolCLASS)

	assert.Equal(t, "ABC", s1.Name())
	assert.Equal(t, "ABC", s1.ScopedName())
	assert.Empty(t, s1.Scope())

	s2 := slang.NewChildSymbol("XYZ", slang.SymbolFIELD, s1)

	assert.Equal(t, "XYZ", s2.Name())
	assert.Equal(t, "ABC.XYZ", s2.ScopedName())
	assert.Equal(t, "ABC", s2.Scope())
}
