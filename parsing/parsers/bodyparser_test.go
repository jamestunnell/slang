package parsers_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/stretchr/testify/assert"
)

func testBodyParserSuccess(
	t *testing.T,
	p parsers.BodyParser,
	input string,
	expected ...*statements.Statement,
) {
	l := lexing.NewLexer(strings.NewReader(input))
	seq := parsing.NewTokenSeq(l)

	if !assert.True(t, p.Run(seq)) {
		for _, err := range p.GetErrors() {
			t.Logf("parse failure: %v", err)
		}
	}

	verifyStatements(t, expected, p.GetStatements())
}

func testBodyParserFail(
	t *testing.T,
	p parsers.BodyParser,
	input string,
) {
	l := lexing.NewLexer(strings.NewReader(input))
	seq := parsing.NewTokenSeq(l)

	p.Run(seq)

	assert.NotEmpty(t, p.GetErrors)
}
