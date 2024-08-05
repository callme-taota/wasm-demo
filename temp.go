package main

import (
	"fmt"
	"github.com/goplus/gop/ast"
	"github.com/goplus/gop/parser"
	"github.com/goplus/gop/token"
	token2 "go/token"
)

func TestGetTokenInfo() {
	src := `
	onStart => {
		flag := true
		for flag {
			onMsg "die", => {
				flag = false
				a
			}
			glide -877, 180, 3
			setXYpos -240, 180
		}
	}
	`

	fset := token.NewFileSet()
	pos := 40
	fset.Position(token2.Pos(pos))

	// 解析源代码
	node, err := parser.ParseFile(fset, "test.spx", src, parser.ParseComments)
	if err != nil {
		fmt.Println(err)
		return
	}

	ast.Inspect(node, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		// 获取节点的位置信息
		nodePos := fset.Position(n.Pos())
		nodeEnd := fset.Position(n.End())

		// 检查光标是否在当前节点范围内
		if pos >= nodePos.Offset && pos <= nodeEnd.Offset {
			// 根据不同节点类型获取不同的信息
			switch x := n.(type) {
			case *ast.FuncDecl:
				fmt.Printf("Function: %s\n", x.Name.Name)
				fmt.Println("Parameters:")
				for _, param := range x.Type.Params.List {
					for _, name := range param.Names {
						fmt.Printf("  %s %s\n", name.Name, param.Type)
					}
				}
				if x.Type.Results != nil {
					fmt.Println("Return types:")
					for _, result := range x.Type.Results.List {
						fmt.Printf("  %s\n", result.Type)
					}
				}
				if x.Doc != nil {
					fmt.Println("Comments:")
					for _, comment := range x.Doc.List {
						fmt.Println(comment.Text)
					}
				}
			case *ast.GenDecl:
				// 变量或常量声明
				for _, spec := range x.Specs {
					switch spec := spec.(type) {
					case *ast.ValueSpec:
						for _, name := range spec.Names {
							fmt.Printf("Name: %s\n", name.Name)
							fmt.Printf("Type: %s\n", spec.Type)
							if len(spec.Values) > 0 {
								fmt.Printf("Value: %s\n", spec.Values[0])
							}
						}
					}
				}
			case *ast.BasicLit:
				// 基本字面量
				fmt.Printf("Literal: %s\n", x.Value)
			case *ast.Ident:
				// 标识符
				fmt.Printf("Identifier: %s\n", x.Name)
				fmt.Println(x.NamePos, x.Obj, x.String())
			// 你可以根据需要添加更多节点类型的处理
			default:
				fmt.Printf("Node: %s\n", fset.Position(n.Pos()))
			}
		}
		return true
	})
}

func TestFunctionComment() {
	// 源代码
	_ = `
	package main

	//aaa
	func rev(f bool) bool {
		return !f
	}

	/*

	@unit: number
	*/
	func add(a int, b int) int {
		return a + b
	}
	`

	spxSrc := `
		glide -877, 180, 3
	`

	fset := token.NewFileSet()

	// 解析源代码
	node, err := parser.ParseFile(fset, "test.spx", spxSrc, parser.ParseComments)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 遍历 AST 并提取函数相关内容
	ast.Inspect(node, func(n ast.Node) bool {
		// 检查节点是否为函数声明
		if fn, ok := n.(*ast.FuncDecl); ok {
			// 打印函数名称
			fmt.Printf("Function name: %s\n", fn.Name.Name)

			// 打印函数参数
			fmt.Println("Parameters:")
			for _, param := range fn.Type.Params.List {
				for _, name := range param.Names {
					fmt.Printf("  %s %s\n", name.Name, param.Type)
				}
			}

			// 打印返回类型
			if fn.Type.Results != nil {
				fmt.Println("Return types:")
				for _, result := range fn.Type.Results.List {
					fmt.Printf("  %s\n", result.Type)
				}
			}

			// 打印函数前的注释
			if fn.Doc != nil {
				fmt.Println("Comments:")
				for _, comment := range fn.Doc.List {
					fmt.Println(comment.Text)
				}
			}
		}
		return true
	})
}

func TestASTree() {
	// 源代码
	src := `
	onStart => {
			flag := true
			for flag {
				onMsg "die", => {
					flag = false
				}
				glide -877, 180, 3
				setXYpos -240, 180
			}
		}
	`

	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, "test.spx", src, parser.AllErrors)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printAstNode(node, fset)
}

func printAstNode(node interface{}, fset *token.FileSet) {
	ast.Print(fset, node)
}

func TestGetTypes() {
	// 源代码
	_ = `
	// 这是一个示例函数
	func add(a int, b int) int {
		if a > b {
			return a
		} else {
			return b
		}
	}

	// 变量定义
	var x int = 10

	// 常量定义
	const y = 20
	`

	src := `onStart => {
			flag := true
			for flag {
				onMsg "die", => {
					flag = false
				}
				glide -877, 180, 3
				setXYpos -240, 180
			}
		}`

	fset := token.NewFileSet()

	// 解析源代码
	node, err := parser.ParseFile(fset, "S.spx", src, parser.ParseComments)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("%+v\\n", node)
	// 遍历 AST 并提取不同节点类型的定义
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			// 打印函数定义
			fmt.Printf("Function name: %s\n", x.Name.Name)
			fmt.Println("Parameters:")
			for _, param := range x.Type.Params.List {
				for _, name := range param.Names {
					fmt.Printf("  %s %s\n", name.Name, param.Type)
				}
			}
			if x.Type.Results != nil {
				fmt.Println("Return types:")
				for _, result := range x.Type.Results.List {
					fmt.Printf("  %s\n", result.Type)
				}
			}
			if x.Doc != nil {
				fmt.Println("Comments:")
				for _, comment := range x.Doc.List {
					fmt.Println(comment.Text)
				}
			}
		case *ast.IfStmt:
			// 打印 if 语句
			fmt.Println("If statement:")
			fmt.Printf("  Condition: %s\n", x.Cond)
			fmt.Printf("  Body: %s\n", x.Body)
			if x.Else != nil {
				fmt.Printf("  Else: %s\n", x.Else)
			}
		case *ast.GenDecl:
			// 打印变量和常量定义
			if x.Tok == token.VAR {
				fmt.Println("Variable declaration:")
				for _, spec := range x.Specs {
					valSpec := spec.(*ast.ValueSpec)
					for _, name := range valSpec.Names {
						fmt.Printf("  %s %s = %s\n", name.Name, valSpec.Type, valSpec.Values)
					}
				}
			} else if x.Tok == token.CONST {
				fmt.Println("Constant declaration:")
				for _, spec := range x.Specs {
					valSpec := spec.(*ast.ValueSpec)
					for _, name := range valSpec.Names {
						fmt.Printf("  %s = %s\n", name.Name, valSpec.Values)
					}
				}
			}
		}
		return true
	})
}

func FindErrors() {
	src := `
	onStart => {
		flag := true
		for flag {
			onMsg "die", => {
				flag = false
				=
			})
			glide -877, 180, 3
			setXYpos -240, 180
		}
	}
	`
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "test.spx", src, parser.AllErrors)

	fmt.Println(f, err)
}

func Suggest() {
	_ = `
	onStart => {
		flag := true
		for flag {
			onMsg "die", => {
				flag = false
				
			}
			glide -877, 180, 3
			setXYpos -240, 180
		}
	}
	`

	//fset := token.NewFileSet()
	//node, err := parser.ParseFile(fset, "filename.spx", src, parser.AllErrors)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	return
	//}
	//
	//prefix := "f"

}

func isLetterOrDigit(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ('0' <= ch && ch <= '9') || ch == '_'
}

type Position struct {
	Line   int    `json:"line"`
	Column int    `json:"column"`
	File   string `json:"file"`
}

type InlayHint struct {
	Type     string   `json:"type"`
	Content  string   `json:"content"`
	Position Position `json:"position"`
}

func GetInlayHint(code string, fileName string) []InlayHint {
	// init hint list
	hints := []InlayHint{}
	// init fset
	fset := token.NewFileSet()
	// init node
	node, err := parser.ParseFile(fset, fileName, code, parser.AllErrors)
	if err != nil {
		hints = append(hints, ErrorParser(err.Error()))
	}
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case (*ast.FuncDecl):
			fmt.Printf("Function name: %s\n", x.Name.Name)
			fmt.Println("Parameters:")
			for _, param := range x.Type.Params.List {
				for _, name := range param.Names {
					fmt.Printf("  %s %s\n", name.Name, param.Type)
				}
			}
			if x.Type.Results != nil {
				fmt.Println("Return types:")
				for _, result := range x.Type.Results.List {
					fmt.Printf("  %s\n", result.Type)
				}
			}
		}
		return true
	})
	return hints
}

func CaseNodeType() {

}

func ErrorParser(e string) InlayHint {
	//TODO(callum-chan): parse error
	if e != "" {
		return InlayHint{
			Type:     "error",
			Content:  e,
			Position: Position{},
		}
	}
	return InlayHint{}
}
