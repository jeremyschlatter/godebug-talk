package main

// This program does code generation. It's taken from the early stages of development of godebug,
// with some modifications. It does two things: (1) add an import statement for "./godebug" and
// (2) insert the function call "godebug.Line()" in lots of places.
//
// As mentioned in the talk, this example does not enter blocks inside of declarations.

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"os"
)

// visitFn is a wrapper to make plain functions implement the ast.Visitor interface.
type visitFn func(ast.Node) ast.Visitor

// Visit is part of the ast.Visitor interface.
func (v visitFn) Visit(n ast.Node) ast.Visitor {
	return v(n)
}

func main() {
	var filename string
	if len(os.Args) != 2 {
		filename = "godebug/longer.go"
	} else {
		filename = os.Args[1]
	}
	// START OMIT
	fs := token.NewFileSet()
	node, err := parser.ParseFile(fs, filename, nil, 0) // HL
	if err != nil {
		log.Fatalf("error during parsing: %v", err)
	}
	ast.Walk(visitFn(addImport), node) // OMIT

	ast.Walk(visitFn(addLineFunc), node) // HL
	// END OMIT

	cfg := printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	cfg.Fprint(os.Stdout, fs, node)
}

func addLineFunc(node ast.Node) ast.Visitor {
	if _, ok := node.(*ast.File); ok {
		return visitFn(addLineFunc)
	}
	return addLineFuncs(node)
}

func addLineFuncs(node ast.Node) ast.Visitor {
	// 2 START OMIT
	if _, ok := node.(*ast.FuncDecl); !ok {
		return nil
	}
	// 2 END OMIT

	fn := node.(*ast.FuncDecl)

	newBody := make([]ast.Stmt, 2*len(fn.Body.List))

	for i := 0; i < len(newBody); i += 2 {
		// 3 START OMIT
		newBody[i] = &ast.ExprStmt{
			X: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("godebug"),
					Sel: ast.NewIdent("Line"),
				},
			},
		}
		// 3 END OMIT
		newBody[i+1] = fn.Body.List[i/2]
	}

	fn.Body.List = newBody
	return nil
}

func addImport(node ast.Node) ast.Visitor {
	if _, ok := node.(*ast.File); ok {
		return visitFn(addImport)
	}
	genDecl, ok := node.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.IMPORT {
		return nil
	}
	genDecl.Specs = append(genDecl.Specs, &ast.ImportSpec{
		Path: &ast.BasicLit{
			Kind:  token.STRING,
			Value: `"./godebug"`,
		},
	})
	return nil
}
