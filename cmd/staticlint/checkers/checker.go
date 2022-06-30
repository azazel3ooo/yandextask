package checkers

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

// ExitCheckAnalyzer - анализатор, который проверяет отсутствие вызова функции  os.Exit() в функции main
// пакета main.
var ExitCheckAnalyzer = &analysis.Analyzer{
	Name: "exitcheck",
	Doc:  "check for calling os.Exit() in main function",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	isMain := func(x *ast.FuncDecl) bool {
		if x.Name.Name == "main" {
			return true
		}
		return false
	}
	checkSelector := func(s *ast.SelectorExpr) {
		isOs := false
		switch x := s.X.(type) {
		case *ast.Ident:
			if x.Name == "os" {
				isOs = true
			}
		}

		if isOs && s.Sel.Name == "Exit" {
			pass.Reportf(s.Sel.NamePos, "calling the os.Exit function")
		}
	}
	checkCallExpr := func(c *ast.CallExpr) {
		switch f := c.Fun.(type) {
		case *ast.SelectorExpr:
			checkSelector(f)
		}
	}
	checkEl := func(c *ast.ExprStmt) {
		switch x := c.X.(type) {
		case *ast.CallExpr:
			checkCallExpr(x)
		}
	}
	checkBody := func(x *ast.BlockStmt) {
		for _, el := range x.List {
			switch c := el.(type) {
			case *ast.ExprStmt:
				checkEl(c)
			}
		}
	}

	for _, file := range pass.Files {
		if file.Name.Name == "main" {
			ast.Inspect(file, func(node ast.Node) bool {
				switch x := node.(type) {
				case *ast.FuncDecl:
					if isMain(x) {
						checkBody(x.Body)
					}
				}
				return true
			})
		}
	}

	return nil, nil
}
