package parsers_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

func TestVarStatementParser(t *testing.T) {
	testVarStmtParser(
		t, "var x 7", statements.NewVar("x", expressions.NewInt(7)))
	testVarStmtParser(
		t, `var name ""`, statements.NewVar("name", expressions.NewStr("")))
}

func testVarStmtParser(
	t *testing.T,
	input string,
	expected *statements.Statement,
) {
	p := parsers.NewVarStatementParser()
	runes := lexing.NewRuneSource(strings.NewReader(input))
	l := lexing.NewLexer(runes)
	seq := parsing.NewTokenSeq(l)

	if !assert.True(t, p.Run(seq, "")) {
		logParseErrs(t, p.GetErrors())

		return
	}

	assert.True(t, p.VarStmt.IsEqual(expected))
}
