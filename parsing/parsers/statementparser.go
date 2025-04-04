package parsers

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/parsing"
)

type StatementParser interface {
	Run(toks slang.TokenSeq, comment string) bool
	GetErrors() []*parsing.ParseErr
	GetStatement() *statements.Statement
}
