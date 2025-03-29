package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type FileParser struct {
	*ParserBase

	Statements []slang.Statement
}

func NewFileParser() *FileParser {
	return &FileParser{
		ParserBase: NewParserBase(),
		Statements: []slang.Statement{},
	}
}

func (p *FileParser) Run(toks slang.TokenSeq) bool {
	p.Statements = []slang.Statement{}

	_ = toks.Skip(slang.TokenNEWLINE)

	for !toks.Current().Is(slang.TokenEOF) {
		if !p.parseStatement(toks) {
			return false
		}

		// if !p.ExpectToken(toks.Current(), slang.TokenNEWLINE, slang.TokenEOF) {
		// 	return false
		// }

		_ = toks.Skip(slang.TokenNEWLINE)
	}

	return true
}

func (p *FileParser) parseStatement(toks slang.TokenSeq) bool {
	commentLines := []string{}

	for toks.Current().Is(slang.TokenCOMMENT) {
		commentLines = append(commentLines, toks.Current().Value())

		if toks.AdvanceSkip(slang.TokenNEWLINE) > 1 || toks.Current().Is(slang.TokenEOF) {
			p.Statements = append(p.Statements, statements.NewComment(commentLines...))

			return true
		}
	}

	stmtParser := p.makeStmtParser(toks.Current())
	if stmtParser == nil {
		return false
	}

	if !p.RunSubParser(toks, stmtParser) {
		return false
	}

	stmt := stmtParser.GetStatement()

	if len(commentLines) > 0 {
		stmt.SetComment(commentLines)
	}

	p.Statements = append(p.Statements, stmt)

	return true
}

func (p *FileParser) makeStmtParser(cur *slang.Token) StatementParser {
	switch cur.Type() {
	case slang.TokenCLASS:
		return NewClassStatementParser()
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
		cur, slang.TokenCONST, slang.TokenCLASS, slang.TokenFUNC, slang.TokenSTRUCT, slang.TokenUSE, slang.TokenVAR)

	return nil
}
