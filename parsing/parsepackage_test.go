package parsing_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jamestunnell/slang/examples"
	"github.com/jamestunnell/slang/parsing"
	"github.com/jamestunnell/slang/parsing/parsers"
)

func TestParsePackag_CalculatorExample(t *testing.T) {
	fp := parsers.NewFileParser()

	modules, err := parsing.ParsePackage(examples.Calculator(), fp)

	assert.NoError(t, err)
	assert.Len(t, modules, 1)
}

func TestParsePackag_GarageExample(t *testing.T) {
	fp := parsers.NewFileParser()

	modules, err := parsing.ParsePackage(examples.Garage(), fp)

	assert.NoError(t, err)
	assert.Len(t, modules, 2)

}
