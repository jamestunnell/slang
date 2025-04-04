package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type VarStatementParser struct {
	*ParserBase

	VarStmt *statements.Statement
}

func NewVarStatementParser() *VarStatementParser {
	return &VarStatementParser{ParserBase: NewParserBase()}
}

func (p *VarStatementParser) GetStatement() *statements.Statement {
	return p.VarStmt
}

func (p *VarStatementParser) Run(toks slang.TokenSeq, comment string) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenVAR) {
		return false
	}

	toks.Advance()

	name, typ, ok := p.ParseNameTypePair(toks)
	if !ok {
		return false
	}

	p.VarStmt = statements.NewVar(name, typ)

	p.VarStmt.SetComment(comment)

	return true
}
