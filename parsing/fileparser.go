package parsing

import (
	"bufio"
	"io"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexer"
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

func ParseFile(r io.Reader) ([]slang.Statement, []*ParseErr) {
	l := lexer.New(bufio.NewReader(r))
	toks := NewTokenSeq(l)
	p := NewFileParser()

	p.Run(toks)

	return p.Statements, p.Errors()
}

func (p *FileParser) Run(toks slang.TokenSeq) {
	p.Statements = []slang.Statement{}

	for !toks.Current().Is(slang.TokenEOF) {
		toks.Skip(slang.TokenNEWLINE)

		if !p.parseStatement(toks) {
			return
		}

		toks.Skip(slang.TokenNEWLINE)
	}
}

func (p *FileParser) parseStatement(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenUSE, slang.TokenFUNC, slang.TokenCLASS) {
		return false
	}

	tokType := toks.Current().Type()

	toks.Advance()

	switch tokType {
	case slang.TokenUSE:
		return p.parseUse(toks)
	case slang.TokenFUNC:
		return p.parseFunc(toks)
	case slang.TokenCLASS:
		return p.parseClass(toks)
	}

	return true
}

func (p *FileParser) parseUse(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenSTRING) {
		return false
	}

	path := toks.Current().Value()

	toks.Advance()

	p.Statements = append(p.Statements, statements.NewUse(path))

	return true
}

func (p *FileParser) parseFunc(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	sigParser := NewFuncSignatureParser()
	if !p.RunSubParser(toks, sigParser) {
		return false
	}

	bodyParser := NewFuncBodyParser()
	if !p.RunSubParser(toks, bodyParser) {
		return false
	}

	fn := ast.NewFunction(
		sigParser.Params, sigParser.ReturnTypes, bodyParser.Statements...)

	p.Statements = append(p.Statements, statements.NewFunc(name, fn))

	return true
}

func (p *FileParser) parseClass(toks slang.TokenSeq) bool {
	if !p.ExpectToken(toks.Current(), slang.TokenSYMBOL) {
		return false
	}

	name := toks.Current().Value()

	toks.Advance()

	classParser := NewClassBodyParser()
	if !p.RunSubParser(toks, classParser) {
		return false
	}

	p.Statements = append(p.Statements, statements.NewClass(name, "", classParser.Statements...))

	return true
}
