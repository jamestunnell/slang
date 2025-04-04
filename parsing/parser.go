package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type Parser interface {
	Run(toks slang.TokenSeq) bool
	GetErrors() []*ParseErr
}

type BodyParser interface {
	Parser

	GetStatements() []*statements.Statement
}

type StatementParser interface {
	Run(toks slang.TokenSeq, comment string) bool
	GetErrors() []*ParseErr
	GetStatement() *statements.Statement
}
