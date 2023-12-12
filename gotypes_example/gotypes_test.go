package gotypes_example_test

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shopspring/decimal"
)

func TestOne(t *testing.T) {

	program := `package main

import "fmt"

func main() {
        fmt.Println("Hello, world")
}`
	confCheckProgram(t, program)
}

func TestTwo(t *testing.T) {

	program := `package mypkg

import (
	"math/big"
)

type MyType struct{
	BI big.Int
}
`
	confCheckProgramWithDefsAndUses(t, "hello.go", program)
}

func Skip_TestThree(t *testing.T) {

	_ = decimal.Decimal{}

	confCheckProgramWithDefsAndUses(t, "somepkg/some.go", nil)
}

type myImporter struct {
	Imp types.Importer
}

func (i myImporter) Import(path string) (*types.Package, error) {

	fmt.Println("myImporter.Import:", path)
	return i.Imp.Import(path)
}

func (i myImporter) ImportFrom(path, dir string, mode types.ImportMode) (*types.Package, error) {

	fmt.Println("myImporter.ImportFrom:", path, dir)
	return i.Imp.(types.ImporterFrom).ImportFrom(path, dir, mode)
}

func confCheckProgram(
	t *testing.T,
	program string,
) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "hello.go", program, 0)
	require.NoError(t, err)

	conf := types.Config{
		Importer:	myImporter{
			Imp: importer.Default(),
		},
	}

	pkg, err := conf.Check(
		"cmd/hello",
		fset,
		[]*ast.File{f},
		nil,
	)
	require.NoError(t, err)

	fmt.Println("pkg.Path:", pkg.Path())
	fmt.Println("pkg.Name:", pkg.Name())
	fmt.Println("pkg.Imports:", pkg.Imports())
	fmt.Println("pkg.Scope:", pkg.Scope())
}

func confCheckProgramWithDefsAndUses(
	t *testing.T,
	filepath string,
	program any,
) {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, filepath, program, 0)
	require.NoError(t, err)

	conf := types.Config{
		Importer:	myImporter{
			Imp: importer.Default(),
		},
	}

	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}

	pkg, err := conf.Check(
		"some/import/hello",
		fset,
		[]*ast.File{f},
		info,
	)
	require.NoError(t, err)

	fmt.Println("pkg.Path:", pkg.Path())
	fmt.Println("pkg.Name:", pkg.Name())
	fmt.Println("pkg.Imports:", pkg.Imports())
	fmt.Println("pkg.Scope:", pkg.Scope())

	fmt.Println("info.Defs:")
	for id, obj := range info.Defs {
		fmt.Printf(
			"Def at %q: %q - %+v\n",
			fset.Position(id.Pos()),
			id.Name,
			obj,
		)
	}

	fmt.Println("info.Uses:")
	for id, obj := range info.Uses {
		fmt.Printf(
			"Use at %q: %q - %+v\n",
			fset.Position(id.Pos()),
			id.Name,
			obj,
		)
	}
}

func Skip_TestImporter(t *testing.T) {

	imp := importer.Default()

	pkg, err := imp.(types.ImporterFrom).ImportFrom("github.com/shopspring/decimal", ".", 0)
	require.NoError(t, err)

	fmt.Println("pkg.Scope:", pkg.Scope().String())

}
