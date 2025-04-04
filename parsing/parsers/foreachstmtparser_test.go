package parsers_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexing"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type forEachStmtParserTest struct {
	Name       string
	Input      string
	ErrorCount int
	ForEach    *statements.Statement
}

func TestForStatementParser(t *testing.T) {
	tests := []*forEachStmtParserTest{
		{
			Name:  "empty",
			Input: "foreach x in y {}",
			ForEach: statements.NewForEach(
				[]string{"x"},
				expressions.NewIdentifier("y"),
				statements.NewBlock(),
			),
		},
		{
			Name:  "two vars, break",
			Input: "foreach a, b in x {break}",
			ForEach: statements.NewForEach(
				[]string{"a", "b"},
				expressions.NewIdentifier("x"),
				statements.NewBlock(statements.NewBreak()),
			),
		},
		{
			Name: "nested if with continue",
			Input: `foreach x in y {
				if x > 2 {
					continue
				}

				fmt.Print("ok")
			}`,
			ForEach: statements.NewForEach(
				[]string{"x"},
				expressions.NewIdentifier("y"),
				statements.NewBlock(
					statements.NewIf(
						expressions.NewGreater(
							expressions.NewIdentifier("x"),
							expressions.NewInt(2),
						),
						statements.NewBlock(statements.NewContinue()),
					),
					statements.NewExpression(
						expressions.NewInvoke(
							expressions.NewAccessMember(expressions.NewIdentifier("fmt"), "Print"),
							expressions.NewInvokeArgPos(expressions.NewStr("ok")),
						),
					),
				),
			),
		},
		{
			Name: "nested foreach",
			Input: `foreach a in b {
				foreach x in y {
					printNums(a + x)
				}
			}`,
			ForEach: statements.NewForEach(
				[]string{"a"},
				expressions.NewIdentifier("b"),
				statements.NewBlock(
					statements.NewForEach(
						[]string{"x"},
						expressions.NewIdentifier("y"),
						statements.NewBlock(
							statements.NewExpression(
								expressions.NewInvoke(
									expressions.NewIdentifier("printNums"),
									expressions.NewInvokeArgPos(
										expressions.NewAdd(
											expressions.NewIdentifier("a"),
											expressions.NewIdentifier("x"),
										),
									),
								),
							),
						),
					),
				),
			),
		},
	}

	for _, test := range tests {
		testForEachStmtParser(t, test)
	}
}

func testForEachStmtParser(t *testing.T, test *forEachStmtParserTest) {
	t.Run(test.Name, func(t *testing.T) {
		p := parsers.NewForEachStmtParser()
		l := lexing.NewLexer(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		assert.True(t, p.Run(seq, ""))

		if !assert.Len(t, p.GetErrors(), test.ErrorCount) {
			logParseErrs(t, p.GetErrors())

			return
		}

		expected, ok := test.ForEach.Core.(*statements.ForEach)

		require.True(t, ok)

		actual, ok := p.Stmt.Core.(*statements.ForEach)

		verifyBlock(t, expected.Block, actual.Block)
	})
}

func verifyBlock(t *testing.T, expected, actual slang.Statement) {
	expectedBlock, ok := expected.(*statements.Statement).Core.(*statements.Block)

	require.True(t, ok)

	actualBlock, ok := actual.(*statements.Statement).Core.(*statements.Block)

	require.True(t, ok)

	verifyStatemnts(t, expectedBlock.Statements, actualBlock.Statements)
}
