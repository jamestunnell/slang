package parsing

import "github.com/jamestunnell/slang"

type Parser interface {
	Run(slang.TokenSeq)
	Errors() []*ParseErr
}
