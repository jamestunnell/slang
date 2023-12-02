package parsing

import "github.com/jamestunnell/slang"

type Parser interface {
	Run(slang.TokenSeq) bool

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
