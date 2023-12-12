package gotypes_example

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/packages"
)

func TestPackages_One(t *testing.T) {

	conf := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			//packages.NeedCompiledGoFiles |
			//packages.NeedImports |
			//packages.NeedTypesSizes |
			packages.NeedDeps |
			packages.NeedTypes,
	}
	pkgs, err := packages.Load(conf, "github.com/shopspring/decimal")
	require.NoError(t, err)

	require.Greater(t, len(pkgs), 0, "no pkgs found")

	fmt.Println("len pkgs", len(pkgs))

	fmt.Println("errors:", packages.PrintErrors(pkgs))

	pkg := pkgs[0]

	fmt.Println("pkg.String:", pkg.String())


	fmt.Println("pkg.Name", pkg.Name)

	fmt.Println("pkg.GoFiles", pkg.GoFiles)

	//fmt.Println("Type errors", len(pkg.TypeErrors))
	//for _, e := range pkg.TypeErrors {
	//	fmt.Println("type error:", e.Error())
	//}

	fmt.Println("pkg.IllTyped", pkg.IllTyped)


	fmt.Println("pkg.Types.Complete:", pkg.Types.Complete())

	//fmt.Println("pkg.Types.Scope:", pkg.Types.Scope().String())

	obj := pkg.Types.Scope().Lookup("Decimal")

	require.NotNil(t, obj, "object not found")

	fmt.Println("obj: ", obj.String())
}

func TestPackages_CurrentModule(t *testing.T) {

	conf := &packages.Config{
		Mode: packages.NeedName |
			packages.NeedFiles |
			//packages.NeedCompiledGoFiles |
			//packages.NeedImports |
			//packages.NeedTypesSizes |
			//packages.NeedDeps |
			packages.NeedModule |
			packages.NeedTypes,
	}
	pkgs, err := packages.Load(conf, "github.com/neurotempest/public_bucket/gotypes_example/somepkg")
	require.NoError(t, err)

	require.Greater(t, len(pkgs), 0, "no pkgs found")

	fmt.Println("len pkgs", len(pkgs))

	fmt.Println("errors:", packages.PrintErrors(pkgs))

	pkg := pkgs[0]

	fmt.Println("pkg.String:", pkg.String())

	fmt.Println("pkg.Name", pkg.Name)

	fmt.Println("pkg.GoFiles", pkg.GoFiles)

	fmt.Println("pkg.Types.Scope:", pkg.Types.Scope().String())

	fmt.Println("pkg.Module:", pkg.Module)
}
