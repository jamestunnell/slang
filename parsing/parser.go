package parsing

import (
	"github.com/jamestunnell/slang/ast/statements"
)

type FileParser interface {
	BodyParser
}

type BodyParser interface {
	Parser

	GetStatements() []*statements.Statement
}

type Parser interface {
	Run(toks TokenSeq) bool
	GetErrors() []*ParseErr
}
