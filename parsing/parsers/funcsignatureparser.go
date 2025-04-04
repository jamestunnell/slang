package parsers

import (
	"github.com/jamestunnell/slang"
)

type FuncSignatureParser struct {
	*ParserBase

	InParams  []slang.Param
	OutParams []slang.Param
}

func NewFuncSignatureParser() *FuncSignatureParser {
	return &FuncSignatureParser{
		ParserBase: NewParserBase(),
		InParams:   []slang.Param{},
		OutParams:  []slang.Param{},
	}
}

func (p *FuncSignatureParser) Run(toks slang.TokenSeq) bool {
	p.InParams = []slang.Param{}
	p.OutParams = []slang.Param{}

	inputSigParser := NewDataSignatureParser()
	if !p.RunSubParser(toks, inputSigParser) {
		return false
	}

	toks.Skip(slang.TokenNEWLINE)

	outParams := []slang.Param{}

	if toks.Current().Is(slang.TokenLPAREN) {
		outputSigParser := NewDataSignatureParser()
		if !p.RunSubParser(toks, outputSigParser) {
			return false
		}

		outParams = outputSigParser.NameTypes
	}

	p.InParams = inputSigParser.NameTypes
	p.OutParams = outParams

	return true

	// if !p.ExpectToken(toks.Current(), slang.TokenLPAREN) {
	// 	return false
	// }

	// toks.AdvanceSkip(slang.TokenNEWLINE)

	// p.Params = []slang.Param{}
	// p.ReturnTypes = []slang.Type{}

	// if !toks.Current().Is(slang.TokenRPAREN) && !p.parseParam(toks) {
	// 	return false
	// }

	// for !toks.Current().Is(slang.TokenRPAREN) {
	// 	if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
	// 		return false
	// 	}

	// 	toks.Advance()

	// 	if !p.parseParam(toks) {
	// 		return false
	// 	}
	// }

	// toks.Advance()

	// // parse return type(s), if any

	// addRetType := func() bool {
	// 	typ, ok := p.ParseType(toks)
	// 	if !ok {
	// 		return false
	// 	}

	// 	p.ReturnTypes = append(p.ReturnTypes, typ)

	// 	return true
	// }

	// switch toks.Current().Type() {
	// case slang.TokenSYMBOL:
	// 	if !addRetType() {
	// 		return false
	// 	}
	// case slang.TokenLPAREN:
	// 	toks.Advance()

	// 	if !addRetType() {
	// 		return false
	// 	}

	// 	for !toks.Current().Is(slang.TokenRPAREN) {
	// 		if !p.ExpectToken(toks.Current(), slang.TokenCOMMA) {
	// 			return false
	// 		}

	// 		toks.Advance()

	// 		if !addRetType() {
	// 			return false
	// 		}
	// 	}

	// 	toks.Advance()
	// }

	// return true
}

// func (p *FuncSignatureParser) parseParam(toks slang.TokenSeq) bool {
// 	names, typ, ok := p.ParseNamesType(toks)
// 	if !ok {
// 		return false
// 	}

// 	for _, name := range names {
// 		p.Params = append(p.Params, ast.NewParam(name, typ))
// 	}

// 	return true
// }
