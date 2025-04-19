package parsers

import (
	"errors"
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type UseStatementParser struct {
	*ParserBase

	UseStmt *statements.Statement
}

var errEmptyUsePath = errors.New("use path is empty")

func NewUseStatementParser() *UseStatementParser {
	return &UseStatementParser{ParserBase: NewParserBase()}
}

func (p *UseStatementParser) GetStatement() *statements.Statement {
	return p.UseStmt
}

func (p *UseStatementParser) Run(
	toks parsing.TokenSeq,
	comment string,
) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenUSE) {
		return false
	}

	toks.Advance()

	var rename string

	if toks.Current().Is(slang.TokenSYMBOL) {
		rename = toks.Current().Value()

		toks.Advance()
	}

	if !p.ExpectToken(toks.Current(), slang.TokenSTRVAL) {
		return false
	}

	pathTok := toks.Current()

	toks.Advance()

	path := pathTok.Value()
	parts := strings.Split(path, "/")

	if len(parts) == 0 {
		parseErr := parsing.NewParseError(errEmptyUsePath, pathTok)

		p.errors = append(p.errors, parseErr)

		return false
	}

	p.UseStmt = statements.NewUse(rename, parts)

	p.UseStmt.SetComment(comment)

	return true
}
