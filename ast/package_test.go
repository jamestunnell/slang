package ast_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/exp/slices"

	"github.com/jamestunnell/slang/ast"
	"github.com/jamestunnell/slang/ast/expressions"
	"github.com/jamestunnell/slang/ast/field"
	"github.com/jamestunnell/slang/ast/statements"
	"github.com/jamestunnell/slang/ast/types"
)

func TestPackage_MarshalUnmarshal(t *testing.T) {
	const pkgName = "garage"

	models := ast.NewModule([]string{"models"},
		statements.NewStruct("Car",
			field.New("Make", types.NewStr()),
			field.New("Model", types.NewStr()),
			field.New("Color", types.NewStr()),
			field.New("VIN", types.NewStr()),
			field.New("Year", types.NewInt()),
			field.New("Miles", types.NewFlt()),
		),
	)
	cars := ast.NewModule([]string{"cars"},
		statements.NewUse("", []string{"..", "modules"}),
		statements.NewFunc("NewFiesta",
			[]*field.Field{
				field.New("vin", types.NewStr()),
				field.New("year", types.NewInt()),
				field.New("color", types.NewStr()),
			},
			[]*field.Field{
				field.New("car", types.NewStruct("models", "Car")),
			},
			statements.NewAssign(
				expressions.NewIdentifier("car"),
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("models"), "Car"),
					expressions.NewInvokeArgKW("Make", expressions.NewStr("Ford")),
					expressions.NewInvokeArgKW("Model", expressions.NewStr("Fiesta")),
					expressions.NewInvokeArgKW("Color", expressions.NewIdentifier("color")),
					expressions.NewInvokeArgKW("Year", expressions.NewIdentifier("year")),
					expressions.NewInvokeArgKW("Miles", expressions.NewFloat(0.0)),
					expressions.NewInvokeArgKW("VIN", expressions.NewIdentifier("vin")),
				),
			),
		),
		statements.NewFunc("NewMustang",
			[]*field.Field{
				field.New("vin", types.NewStr()),
				field.New("year", types.NewInt()),
				field.New("color", types.NewStr()),
			},
			[]*field.Field{
				field.New("car", types.NewStruct("models", "Car")),
			},
			statements.NewAssign(
				expressions.NewIdentifier("car"),
				expressions.NewInvoke(
					expressions.NewAccessMember(expressions.NewIdentifier("models"), "Car"),
					expressions.NewInvokeArgKW("Make", expressions.NewStr("Ford")),
					expressions.NewInvokeArgKW("Model", expressions.NewStr("Mustang")),
					expressions.NewInvokeArgKW("Color", expressions.NewIdentifier("color")),
					expressions.NewInvokeArgKW("Year", expressions.NewIdentifier("year")),
					expressions.NewInvokeArgKW("Miles", expressions.NewFloat(0.0)),
					expressions.NewInvokeArgKW("VIN", expressions.NewIdentifier("vin")),
				),
			),
		),
	)
	pkg := ast.NewPackage(pkgName, cars, models)

	d, err := json.Marshal(pkg)

	require.NoError(t, err)

	var pkg2 ast.Package

	require.NoError(t, json.Unmarshal(d, &pkg2))

	assert.Equal(t, pkg2.PackageInfo.GetName(), pkgName)
	assert.True(t, slices.EqualFunc(pkg.Modules, pkg2.Modules, modulesEqual))
}

func modulesEqual(a, b *ast.Module) bool {
	return a.IsEqual(b)
}
