package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type VarStatementParser struct {
	*ParserBase

	VarStmt *statements.Var
}

func NewVarStatementParser() *VarStatementParser {
	return &VarStatementParser{ParserBase: NewParserBase()}
}

func (p *VarStatementParser) GetStatement() slang.Statement {
	return p.VarStmt
}

func (p *VarStatementParser) Run(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenVAR) {
		return false
	}

	toks.Advance()

	name, typ, ok := p.ParseNameTypePair(toks)
	if !ok {
		return false
	}

	p.VarStmt = statements.NewVar(name, typ)

	return true
}
