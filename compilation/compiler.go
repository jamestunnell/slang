package compilation

type Compiler interface {
	FirstPass() error

	CollectSymbols() []Symbol
}

type Visitor interface {
	Visit(name string, child Compiler)
}
