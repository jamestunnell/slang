package compilation

type MethodCompiler struct {
	*FuncCompiler
}

func NewMethodCompiler(scope string, stmts StmtSeq, parent Compiler) *MethodCompiler {
	return &MethodCompiler{
		FuncCompiler: NewFuncCompiler(scope, stmts, parent),
	}
}
