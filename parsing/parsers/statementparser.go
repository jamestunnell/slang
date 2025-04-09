package parsers

import (
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type StatementParser interface {
	Run(toks parsing.TokenSeq, comment string) bool
	GetErrors() []*parsing.ParseErr
	GetStatement() *statements.Statement
}
