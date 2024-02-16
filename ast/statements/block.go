package statements

import "github.com/jamestunnell/slang"

type Block struct {
	*Base

	Statements []slang.Statement
}

func NewBlock(stmts ...slang.Statement) *Block {
	return &Block{
		Base:       NewBase(slang.StatementBLOCK),
		Statements: stmts,
	}
}

func (b *Block) IsEqual(other slang.Statement) bool {
	b2, ok := other.(*Block)
	if !ok {
		return false
	}

	return slang.StatementsEqual(b.Statements, b2.Statements)
}

func (b *Block) Eval(env slang.Environment) (slang.Objects, error) {
	for _, st := range b.Statements {
		// TODO: how to detect a nested break/continue/return statement?

		results, err := st.Eval(env)
		if err != nil {
			return slang.Objects{}, err
		}

		if len(results) > 0 {
			return results, nil
		}
	}

	return slang.Objects{}, nil
}
