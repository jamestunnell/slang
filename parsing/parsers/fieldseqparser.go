package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/customerrs"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/sliceutil"
)

type FieldSeqParser struct {
	*bodyParser
}

func NewFieldSeqParser() *FieldSeqParser {
	bodyParser := NewBodyParser(
		func(toks parsing.TokenSeq) error {
			if !toks.Current().Is(slang.TokenLPAREN) {
				return customerrs.NewErrWrongTokenType(toks.Current(), slang.TokenLPAREN)
			}

			toks.Advance()

			return nil
		},
		slang.TokenRPAREN,
		func(cur *slang.Token) (StatementParser, error) { return NewFieldStatementParser(), nil },
	)

	return &FieldSeqParser{bodyParser: bodyParser}
}

func (p *FieldSeqParser) GetFields() []*field.Field {
	fields := []*field.Field{}

	fieldStmts := sliceutil.Where(p.GetStatements(),
		func(s *statements.Statement) bool { return s.Type == slang.StatementFIELD })

	for _, stmt := range fieldStmts {
		core := stmt.Core.(*statements.Field)
		moreFields := sliceutil.Map(core.Names, func(name string) *field.Field {
			return field.New(name, core.Type)
		})

		fields = append(fields, moreFields...)
	}

	return fields
}
