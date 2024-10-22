package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/lexing"
)

type Parser interface {
	Run(lexing.TokenSeq) bool

	GetErrors() []*ParseErr
}

type BodyParser interface {
	Parser

	GetStatements() []slang.Statement
}

type StatementParser interface {
	Parser

	GetStatement() slang.Statement
}
