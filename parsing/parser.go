package parsing

import "github.com/jamestunnell/slang"

type Parser interface {
	Run(slang.TokenSeq)

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
