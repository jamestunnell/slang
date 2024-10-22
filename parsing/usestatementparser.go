package parsing

import (
	"errors"
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
)

type UseStatementParser struct {
	*ParserBase

	UseStmt *statements.Use
}

var errEmptyUsePath = errors.New("use path is empty")

func NewUseStatementParser() *UseStatementParser {
	return &UseStatementParser{ParserBase: NewParserBase()}
}

func (p *UseStatementParser) GetStatement() slang.Statement {
	return p.UseStmt
}

func (p *UseStatementParser) Run(toks lexing.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), lexing.TokenUSE) {
		return false
	}

	toks.Advance()

	if !p.ExpectToken(toks.Current(), lexing.TokenSTRING) {
		return false
	}

	pathTok := toks.Current()

	toks.Advance()

	path := pathTok.Value()
	parts := strings.Split(path, "/")

	if len(parts) == 0 {
		parseErr := NewParseError(errEmptyUsePath, pathTok)

		p.errors = append(p.errors, parseErr)

		return false
	}

	p.UseStmt = statements.NewUse(parts...)

	return true
}
