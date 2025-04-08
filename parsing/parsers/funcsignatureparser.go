package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
)

type FuncSignatureParser struct {
	*ParserBase

	Inputs  []*field.Field
	Outputs []*field.Field
}

func NewFuncSignatureParser() *FuncSignatureParser {
	return &FuncSignatureParser{
		ParserBase: NewParserBase(),
		Inputs:     []*field.Field{},
		Outputs:    []*field.Field{},
	}
}

func (p *FuncSignatureParser) Run(toks slang.TokenSeq) bool {
	p.Inputs = []*field.Field{}
	p.Outputs = []*field.Field{}

	inputsParser := NewFieldSeqParser()
	if !p.RunSubParser(toks, inputsParser) {
		return false
	}

	toks.Skip(slang.TokenNEWLINE)

	outputs := []*field.Field{}

	if toks.Current().Is(slang.TokenLPAREN) {
		outputsParser := NewFieldSeqParser()
		if !p.RunSubParser(toks, outputsParser) {
			return false
		}

		outputs = outputsParser.GetFields()
	}

	p.Inputs = inputsParser.GetFields()
	p.Outputs = outputs

	return true

	// if !p.ExpectToken(toks.Current(), slang.TokenLPAREN) {
	// 	return false
	// }

	// toks.AdvanceSkip(slang.TokenNEWLINE)

	// p.Params = []*types.NameType{}
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
