package parsing_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/lexer"
	"github.com/jamestunnell/slang/parsing"
)

type classParserSuccessTest struct {
	TestName string
	Input    string
	Fields   map[string]string
	Methods  map[string]*ast.Function
}

func TestClassBodyParserFailure(t *testing.T) {
	testClassBodyParserFail(t, "no input", "")
}

func TestClassBodyParserSuccess(t *testing.T) {
	tests := []*classParserSuccessTest{
		{
			TestName: "empty",
			Input:    `{}`,
			Fields:   map[string]string{},
			Methods:  map[string]*ast.Function{},
		},
		{
			TestName: "with fields",
			Input: `{
				field X myMod.myType
				field Y float
			}`,
			Fields: map[string]string{
				"X": "myMod.myType",
				"Y": "float",
			},
			Methods: map[string]*ast.Function{},
		},
		{
			TestName: "with empty methods",
			Input: `{
				method X(){}
				method Y(a int){}
			}`,
			Fields: map[string]string{},
			Methods: map[string]*ast.Function{
				"X": ast.NewFunction([]*ast.Param{}, []string{}),
				"Y": ast.NewFunction([]*ast.Param{ast.NewParam("a", "int")}, []string{}),
			},
		},
		{
			TestName: "non-empty method",
			Input: `{
				method sub(x int, y int) int {
					return x - y
				}
			}`,
			Fields: map[string]string{},
			Methods: map[string]*ast.Function{
				"sub": ast.NewFunction(
					[]*ast.Param{{Name: "x", Type: "int"}, {Name: "y", Type: "int"}},
					[]string{"int"},
					statements.NewReturn(sub(id("x"), id("y"))),
				),
			},
		},
	}

	for _, test := range tests {
		testClassBodyParserSuccess(t, test)
	}
}

func testClassBodyParserSuccess(t *testing.T, test *classParserSuccessTest) {
	t.Run(test.TestName, func(t *testing.T) {
		p := parsing.NewClassBodyParser()
		l := lexer.New(strings.NewReader(test.Input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		if !assert.Empty(t, p.Errors()) {
			for _, parseErr := range p.Errors() {
				t.Logf("unxpected parse err at %s: %v", parseErr.Token.Location, parseErr.Error)
			}

			t.FailNow()
		}

		cls := p.Class()

		require.NotNil(t, cls)
		require.ElementsMatch(t, maps.Keys(test.Fields), cls.GetFieldNames())
		require.ElementsMatch(t, maps.Keys(test.Methods), cls.GetMethodNames())

		for name, expectedType := range test.Fields {
			actual, ok := cls.GetFieldType(name)

			if !assert.True(t, ok) {
				t.Logf("no type found for field %s", name)

				continue
			}

			assert.Equal(t, expectedType, actual)
		}

		for name, expectedMethod := range test.Methods {
			actual, ok := cls.GetMethod(name)

			if !assert.True(t, ok) {
				t.Logf("method %s not found", name)

				continue
			}

			verifyFunc(t, expectedMethod, actual)
		}
	})
}

func testClassBodyParserFail(t *testing.T, testName, input string) {
	t.Run(testName, func(t *testing.T) {
		p := parsing.NewClassBodyParser()
		l := lexer.New(strings.NewReader(input))
		seq := parsing.NewTokenSeq(l)

		p.Run(seq)

		assert.NotEmpty(t, p.Errors)
	})
}
