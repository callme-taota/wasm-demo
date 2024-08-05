package utils

import (
	"fmt"
	"go/types"

	"github.com/goplus/gop"
	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	"github.com/goplus/gop/x/typesutil"
	"github.com/goplus/mod/env"
	"github.com/goplus/mod/gopmod"
	"github.com/goplus/mod/modfile"
	"github.com/goplus/mod/modload"
)

var spxProject = &modfile.Project{
	Ext: ".gmx", Class: "*Game",
	Works:    []*modfile.Class{{Ext: ".spx", Class: "Sprite"}},
	PkgPaths: []string{"github.com/goplus/spx", "math"}}

func StartSPXTypesAnalyser(fileName string, fileCode string) interface{} {
	// init fset
	fileSet := token.NewFileSet()
	// init spx mode
	mod := initSPXMod()
	// init conf
	conf := initSPXParserConf()

	info, err := spxInfo(mod, fileSet, fileName, fileCode, conf)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	// convert type info to some valid value
	defs := ""
	for k, v := range info.Defs {
		defs += fmt.Sprintf("k: %v, v: %v\n", k, v)
	}
	types := ""
	for k, v := range info.Types {
		types += fmt.Sprintf("k: %v, v: %v\n", k, v)
	}
	instances := ""
	for k, v := range info.Instances {
		instances += fmt.Sprintf("k: %v, v: %v\n", k, v)
	}
	result := map[string]interface{}{
		"Defs":      defs,
		"Types":     types,
		"Instances": instances,
	}

	return result
}

// init function
func initSPXMod() *gopmod.Module {
	//init spxMod
	var spxMod *gopmod.Module
	spxMod = gopmod.New(modload.Default)
	spxMod.Opt.Projects = append(spxMod.Opt.Projects, spxProject)
	spxMod.ImportClasses()
	return spxMod
}

// init function
func initSPXParserConf() parser.Config {
	return parser.Config{
		ClassKind: func(fname string) (isProj bool, ok bool) {
			ext := modfile.ClassExt(fname)
			c, ok := lookupClass(ext)
			if ok {
				isProj = c.IsProj(ext, fname)
			}
			return
		},
	}
}

// check function
func lookupClass(ext string) (c *modfile.Project, ok bool) {
	switch ext {
	case ".gmx", ".spx":
		return spxProject, true
	}
	return
}

func spxInfo(mod *gopmod.Module, fileSet *token.FileSet, fileName string, fileCode string, parseConf parser.Config) (*typesutil.Info, error) {
	// new parser
	file, err := parser.ParseEntry(fileSet, fileName, fileCode, parseConf)
	if err != nil {
		return nil, err
	}

	// init types conf
	conf := &types.Config{}
	conf.Importer = gop.NewImporter(nil, &env.Gop{Root: "../..", Version: "1.0"}, fileSet)
	chkOpts := &typesutil.Config{
		Types:                 types.NewPackage("main", file.Name.Name),
		Fset:                  fileSet,
		Mod:                   mod,
		UpdateGoTypesOverload: false,
	}

	// init info
	info := &typesutil.Info{
		Types:      make(map[ast.Expr]types.TypeAndValue),
		Defs:       make(map[*ast.Ident]types.Object),
		Uses:       make(map[*ast.Ident]types.Object),
		Implicits:  make(map[ast.Node]types.Object),
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
		Scopes:     make(map[ast.Node]*types.Scope),
		Overloads:  make(map[*ast.Ident][]types.Object),
	}
	check := typesutil.NewChecker(conf, chkOpts, nil, info)
	err = check.Files(nil, []*ast.File{file})
	return info, nil
}
