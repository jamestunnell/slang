package parsing

import "github.com/jamestunnell/slang"

type Parser interface {
	Run(slang.TokenSeq)

	GetErrors() []*ParseErr
}

type StatementParser interface {
	Parser

	GetStatements() []slang.Statement
}
