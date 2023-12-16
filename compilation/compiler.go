package compilation

import "github.com/jamestunnell/slang"

type Compiler interface {
	FirstPass() error

	Symbol() *slang.Symbol
	ChildSymbols() []*slang.Symbol
}

type Visitor interface {
	Visit(name string, child Compiler)
}
