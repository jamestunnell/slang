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
	runes := lexing.NewRuneSource(strings.NewReader(input))
	l := lexing.NewLexer(runes)
	toks := parsing.NewTokenSeq(l)

	if !assert.True(t, p.Run(toks)) {
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
	p.Run(tokenSeqFromStr(input))

	assert.NotEmpty(t, p.GetErrors)
}

func tokenSeqFromStr(input string) parsing.TokenSeq {
	runes := lexing.NewRuneSource(strings.NewReader(input))
	l := lexing.NewLexer(runes)

	return parsing.NewTokenSeq(l)
}
