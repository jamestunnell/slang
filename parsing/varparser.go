package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type VarParser struct {
	*ParserBase

	VarStmt *statements.Var
}

func NewVarParser() *VarParser {
	return &VarParser{ParserBase: NewParserBase()}
}

func (p *VarParser) GetStatement() slang.Statement {
	return p.VarStmt
}

func (p *VarParser) Run(toks slang.TokenSeq) bool {
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
