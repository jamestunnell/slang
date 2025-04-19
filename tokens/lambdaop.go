package tokens

import "github.com/jamestunnell/slang"

type LambdaOp struct{}

const StrLAMBDAOP = "=>"

func LAMBDAOP() slang.TokenInfo           { return &LambdaOp{} }
func (t *LambdaOp) Type() slang.TokenType { return slang.TokenMINUSMINUS }
func (t *LambdaOp) Value() string         { return StrLAMBDAOP }
