package parsing_test

import (
	"strings"
	"testing"

	"github.com/jamestunnell/slang"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type forEachStmtParserTest struct {
	Name       string
	Input      string
	ErrorCount int
	ForEach    *statements.ForEach
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
				statements.NewBlock(
					statements.NewBreak(),
				),
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
							expressions.NewInteger(2),
						),
						statements.NewBlock(
							statements.NewContinue(),
						),
					),
					statements.NewExpression(
						expressions.NewCall(
							expressions.NewAccessMember(expressions.NewIdentifier("fmt"), "Print"),
							expressions.NewPositionalArg(expressions.NewString("ok")),
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
								expressions.NewCall(
									expressions.NewIdentifier("printNums"),
									expressions.NewPositionalArg(
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
		p := parsing.NewForEachStmtParser()
		l := lexer.New(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		assert.True(t, p.Run(seq))

		if !assert.Len(t, p.GetErrors(), test.ErrorCount) {
			logParseErrs(t, p.GetErrors())

			return
		}

		actual, ok := p.Stmt.(*statements.ForEach)

		require.True(t, ok)

		verifyBlock(t, test.ForEach.Block, actual.Block)
	})
}

func verifyBlock(t *testing.T, expected, actual slang.Statement) {
	expectedBlock, ok := expected.(*statements.Block)

	require.True(t, ok)

	actualBlock, ok := actual.(*statements.Block)

	require.True(t, ok)

	verifyStatemnts(t, expectedBlock.Statements, actualBlock.Statements)
}
