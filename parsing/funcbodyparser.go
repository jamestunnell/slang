package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

type FuncBodyParser struct {
	*BodyParserBase
}

func NewFuncBodyParser() *FuncBodyParser {
	p := &FuncBodyParser{}

	p.BodyParserBase = NewBodyParserBase(p.parseStatement)

	return p
}

func (p *FuncBodyParser) parseStatement(toks lexing.TokenSeq) slang.Statement {
	var sp StatementParser

	switch toks.Current().Type {
	case lexing.TokenCONST:
		sp = NewConstStatementParser()
	case lexing.TokenIF:
		sp = NewIfStatementParser()
	case lexing.TokenFOREACH:
		sp = NewForEachStmtParser()
	case lexing.TokenRETURN:
		sp = NewReturnStatementParser()
	case lexing.TokenVAR:
		sp = NewVarStatementParser()
	default:
		sp = NewExprOrAssignStatementParser()
	}

	return p.ParseStatement(toks, sp)
}
