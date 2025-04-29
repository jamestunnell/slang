package app

import (
	"strings"

	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

func parseInput(input string) ([]*statements.Statement, error) {
	runes := lexing.NewRuneSource(strings.NewReader(input))
	l := lexing.NewLexer(runes)
	p := parsers.NewFileParser()

	if err := parsing.RunParser(l, p); err != nil {
		return []*statements.Statement{}, err
	}

	return p.GetStatements(), nil
}
