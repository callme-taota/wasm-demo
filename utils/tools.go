package utils

import (
	"embed"
	"fmt"
	"go/types"
	"strings"
	"unicode"

	"github.com/goplus/gop/token"
	"golang.org/x/tools/go/gcexportdata"
	"golang.org/x/tools/go/types/typeutil"
)

var pkg *types.Package

func getSPXTypes(spx embed.FS, fset *token.FileSet) *types.Package {
	f, err := spx.Open("spx.a")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer f.Close()
	r, err := gcexportdata.NewReader(f)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	const primary = "<primary>"
	imports := make(map[string]*types.Package)
	pkg, err = gcexportdata.Read(r, fset, imports, primary)

	return pkg
}

func findFromSPXTypes(spx embed.FS, fset *token.FileSet, target string) *types.Selection {
	if pkg == nil {
		getSPXTypes(spx, fset)
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		// For types, print each method.
		if _, ok := obj.(*types.TypeName); ok {
			for _, method := range typeutil.IntuitiveMethodSet(obj.Type(), nil) {
				//fmt.Printf("%s: %s\n",
				//	fset.Position(method.Obj().Pos()),
				//	types.SelectionString(method, qual))
				if checkStringMatch(method.Obj().Name(), target) {
					return method
				}

			}
		}
	}
	return nil
}

func checkStringMatch(s, target string) bool {
	target = capitalizeFirstLetter(target)
	if s == target {
		return true
	}
	return strings.Contains(s, target) && strings.Contains(s, "__")
}

func capitalizeFirstLetter(s string) string {
	if s == "" {
		return s
	}

	// Convert string to []rune to handle Unicode characters correctly.
	runes := []rune(s)

	// Capitalize the first character if it's a letter.
	if unicode.IsLower(runes[0]) {
		runes[0] = unicode.ToUpper(runes[0])
	}

	// Convert []rune back to string and return it.
	return string(runes)
}
