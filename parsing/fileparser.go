package parsing

import (
	"strings"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type FileParser struct {
	*ParserBase

	Statements []*statements.Statement
}

func NewFileParser() *FileParser {
	return &FileParser{
		ParserBase: NewParserBase(),
		Statements: []*statements.Statement{},
	}
}

func (p *FileParser) Run(toks slang.TokenSeq) bool {
	p.Statements = []*statements.Statement{}

	_ = toks.Skip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenEOF) {
		if !p.parseStatement(toks) {
			return false
		}

		_ = toks.Skip(slang.TokenNEWLINE)
	}

	return true
}

func (p *FileParser) parseStatement(toks slang.TokenSeq) bool {
	commentLines := []string{}

	for toks.Current().Is(slang.TokenCOMMENT) {
		commentLines = append(commentLines, toks.Current().Value())

		if toks.AdvanceSkip(slang.TokenNEWLINE) > 1 || toks.Current().Is(slang.TokenEOF) {
			stmt := statements.NewComment()

			stmt.SetComment(makeComment(commentLines))

			p.Statements = append(p.Statements, stmt)

			return true
		}
	}

	stmtParser := p.makeStmtParser(toks.Current())
	if stmtParser == nil {
		return false
	}

	if !p.RunSubStmtParser(toks, makeComment(commentLines), stmtParser) {
		return false
	}

	p.Statements = append(p.Statements, stmtParser.GetStatement())

	return true
}

func (p *FileParser) makeStmtParser(cur *slang.Token) StatementParser {
	switch cur.Type() {
	case slang.TokenCONST:
		return NewConstStatementParser()
	case slang.TokenFUNC:
		return NewFuncStatementParser()
	case slang.TokenSTRUCT:
		return NewStructStatementParser()
	case slang.TokenVAR:
		return NewVarStatementParser()
	case slang.TokenUSE:
		return NewUseStatementParser()
	}

	p.TokenErr(
		cur, slang.TokenCONST, slang.TokenFUNC, slang.TokenSTRUCT, slang.TokenUSE, slang.TokenVAR)

	return nil
}

func makeComment(lines []string) string {
	return strings.Join(lines, " ")
}
