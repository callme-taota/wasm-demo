package utils

//
//import (
//	"bytes"
//	"fmt"
//	goast "go/ast"
//	"go/types"
//	"path/filepath"
//
//	"github.com/goplus/gogen"
//	"github.com/goplus/gop/ast"
//	"github.com/goplus/gop/cl"
//	"github.com/goplus/gop/parser"
//	"github.com/goplus/gop/token"
//	"github.com/goplus/igop"
//	"github.com/goplus/mod/modfile"
//
//	_ "github.com/goplus/igop/pkg/bufio"
//	_ "github.com/goplus/igop/pkg/context"
//	_ "github.com/goplus/igop/pkg/errors"
//	_ "github.com/goplus/igop/pkg/fmt"
//	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin"
//	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/iox"
//	_ "github.com/goplus/igop/pkg/github.com/goplus/gop/builtin/ng"
//	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/errors"
//	_ "github.com/goplus/igop/pkg/github.com/qiniu/x/gsh"
//	_ "github.com/goplus/igop/pkg/io"
//	_ "github.com/goplus/igop/pkg/log"
//	_ "github.com/goplus/igop/pkg/math"
//	_ "github.com/goplus/igop/pkg/math/big"
//	_ "github.com/goplus/igop/pkg/math/bits"
//	_ "github.com/goplus/igop/pkg/os"
//	_ "github.com/goplus/igop/pkg/os/exec"
//	_ "github.com/goplus/igop/pkg/path/filepath"
//	_ "github.com/goplus/igop/pkg/runtime"
//	_ "github.com/goplus/igop/pkg/strconv"
//	_ "github.com/goplus/igop/pkg/strings"
//)
//
//type Class = cl.Class
//
//var (
//	projects = make(map[string]*cl.Project)
//)
//
//func RegisterClassFileType(ext string, class string, works []*Class, pkgPaths ...string) {
//	cls := &cl.Project{
//		Ext:      ext,
//		Class:    class,
//		Works:    works,
//		PkgPaths: pkgPaths,
//	}
//	if ext != "" {
//		projects[ext] = cls
//	}
//	for _, w := range works {
//		projects[w.Ext] = cls
//	}
//}
//
//func init() {
//	cl.SetDebug(cl.FlagNoMarkAutogen)
//	igop.RegisterFileProcess(".gop", BuildFile)
//	igop.RegisterFileProcess(".gox", BuildFile)
//	igop.RegisterFileProcess(".gsh", BuildFile)
//	RegisterClassFileType(".gmx", "Game", []*Class{{Ext: ".spx", Class: "Sprite"}}, "github.com/goplus/spx", "math")
//	RegisterClassFileType(".gsh", "App", nil, "github.com/qiniu/x/gsh", "math")
//}
//
//func GetPkg(filename string, src interface{}) (*Package, error) {
//	ctx := igop.NewContext(0)
//	c := NewContext(ctx)
//	pkg, err := c.ParseFile(filename, src)
//	return pkg, err
//}
//
//func BuildFile(ctx *igop.Context, filename string, src interface{}) (data []byte, err error) {
//	defer func() {
//		r := recover()
//		if r != nil {
//			err = fmt.Errorf("compile %v failed. %v", filename, r)
//		}
//	}()
//	c := NewContext(ctx)
//	pkg, err := c.ParseFile(filename, src)
//	if err != nil {
//		return nil, err
//	}
//	return pkg.ToSource()
//}
//
//func BuildFSDir(ctx *igop.Context, fs parser.FileSystem, dir string) (data []byte, err error) {
//	defer func() {
//		r := recover()
//		if r != nil {
//			err = fmt.Errorf("compile %v failed. %v", dir, err)
//		}
//	}()
//	c := NewContext(ctx)
//	pkg, err := c.ParseFSDir(fs, dir)
//	if err != nil {
//		return nil, err
//	}
//	return pkg.ToSource()
//}
//
//func BuildDir(ctx *igop.Context, dir string) (data []byte, err error) {
//	defer func() {
//		r := recover()
//		if r != nil {
//			err = fmt.Errorf("compile %v failed. %v", dir, err)
//		}
//	}()
//	c := NewContext(ctx)
//	pkg, err := c.ParseDir(dir)
//	if err != nil {
//		return nil, err
//	}
//	return pkg.ToSource()
//}
//
//type Package struct {
//	Fset *token.FileSet
//	Pkg  *gogen.Package
//}
//
//func (p *Package) ToSource() ([]byte, error) {
//	var buf bytes.Buffer
//	if err := p.Pkg.WriteTo(&buf); err != nil {
//		return nil, err
//	}
//	return buf.Bytes(), nil
//}
//
//func (p *Package) ToAst() *goast.File {
//	return p.Pkg.ASTFile()
//}
//
//type Context struct {
//	ctx  *igop.Context
//	fset *token.FileSet
//	imp  *igop.Importer
//	gop  igop.Loader
//}
//
//func ClassKind(fname string) (isProj, ok bool) {
//	ext := modfile.ClassExt(fname)
//	switch ext {
//	case ".gmx", ".gsh":
//		return true, true
//	case ".spx":
//		return fname == "main.spx", true
//	default:
//		if c, ok := projects[ext]; ok {
//			for _, w := range c.Works {
//				if w.Ext == ext {
//					if ext != c.Ext || fname != "main"+ext {
//						return false, true
//					}
//					break
//				}
//			}
//			return true, true
//		}
//	}
//	return
//}
//
//func NewContext(ctx *igop.Context) *Context {
//	if ctx.IsEvalMode() {
//		ctx = igop.NewContext(0)
//	}
//	ctx.Mode |= igop.CheckGopOverloadFunc
//	return &Context{ctx: ctx, imp: igop.NewImporter(ctx), fset: token.NewFileSet(), gop: igop.NewTypesLoader(ctx, 0)}
//}
//
//func isGopPackage(path string) bool {
//	if pkg, ok := igop.LookupPackage(path); ok {
//		if _, ok := pkg.UntypedConsts["GopPackage"]; ok {
//			return true
//		}
//	}
//	return false
//}
//
//func (c *Context) Import(path string) (*types.Package, error) {
//	if isGopPackage(path) {
//		return c.gop.Import(path)
//	}
//	return c.imp.Import(path)
//}
//
//func (c *Context) ParseDir(dir string) (*Package, error) {
//	pkgs, err := parser.ParseDirEx(c.fset, dir, parser.Config{
//		ClassKind: ClassKind,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return c.loadPackage(dir, pkgs)
//}
//
//func (c *Context) ParseFSDir(fs parser.FileSystem, dir string) (*Package, error) {
//	pkgs, err := parser.ParseFSDir(c.fset, fs, dir, parser.Config{
//		ClassKind: ClassKind,
//	})
//	if err != nil {
//		return nil, err
//	}
//	return c.loadPackage(dir, pkgs)
//}
//
//func (c *Context) ParseFile(fname string, src interface{}) (*Package, error) {
//	srcDir, _ := filepath.Split(fname)
//	fnameRmGox := fname
//	ext := filepath.Ext(fname)
//	var err error
//	var isProj, isClass, isNormalGox, rmGox bool
//	switch ext {
//	case ".go", ".gop":
//	case ".gox":
//		isClass = true
//		t := fname[:len(fname)-4]
//		if c := filepath.Ext(t); c != "" {
//			ext, fnameRmGox, rmGox = c, t, true
//		} else {
//			isNormalGox = true
//		}
//		fallthrough
//	default:
//		if !isNormalGox {
//			if isProj, isClass = ClassKind(fnameRmGox); !isClass {
//				if rmGox {
//					return nil, fmt.Errorf("not found Go+ class by ext %q for %q", ext, fname)
//				}
//				return nil, nil
//			}
//		}
//	}
//	if err != nil {
//		return nil, err
//	}
//	mode := parser.ParseComments
//	if isClass {
//		mode |= parser.ParseGoPlusClass
//	}
//	f, err := parser.ParseFile(c.fset, fname, src, mode)
//	if err != nil {
//		return nil, err
//	}
//	f.IsProj, f.IsClass, f.IsNormalGox = isProj, isClass, isNormalGox
//	name := f.Name.Name
//	pkgs := map[string]*ast.Package{
//		name: &ast.Package{
//			Name: name,
//			Files: map[string]*ast.File{
//				fname: f,
//			},
//		},
//	}
//	return c.loadPackage(srcDir, pkgs)
//}
//
//func (c *Context) loadPackage(srcDir string, pkgs map[string]*ast.Package) (*Package, error) {
//	mainPkg, ok := pkgs["main"]
//	if !ok {
//		for _, v := range pkgs {
//			mainPkg = v
//			break
//		}
//	}
//	if c.ctx.Mode&igop.DisableCustomBuiltin == 0 {
//		if f, err := igop.ParseBuiltin(c.fset, mainPkg.Name); err == nil {
//			mainPkg.GoFiles = map[string]*goast.File{"_igop_builtin.go": f}
//		}
//	}
//	conf := &cl.Config{Fset: c.fset}
//	conf.Importer = c
//	conf.LookupClass = func(ext string) (c *cl.Project, ok bool) {
//		c, ok = projects[ext]
//		return
//	}
//	if c.ctx.IsEvalMode() {
//		conf.NoSkipConstant = true
//	}
//	out, err := cl.NewPackage("", mainPkg, conf)
//	if err != nil {
//		return nil, err
//	}
//	return &Package{c.fset, out}, nil
//}
