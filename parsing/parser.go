package parsing

import (
	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/statements"
)

type FileParser interface {
	Parser

	GetStatements() []*statements.Statement
}

type Parser interface {
	Run(toks slang.TokenSeq) bool
	GetErrors() []*ParseErr
}
