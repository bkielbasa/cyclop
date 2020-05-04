package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

type analyzer struct {
	noTests bool
	top     int
}

func (a analyzer) analyze(paths []string) ([]stat, error) {
	stats := []stat{}

	for _, p := range paths {
		if isDir(p) {
			s, err := a.analyzeDir(p)
			if err != nil {
				return nil, err
			}

			stats = append(stats, s...)
		} else {
			r, err := a.fileReport(p)
			if err != nil {
				return nil, err
			}

			stats = append(stats, r...)
		}
	}

	return stats, nil
}

func (a analyzer) analyzeDir(path string) ([]stat, error) {
	stats := []stat{}

	filepath.Walk(path, func(p string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			r, e := a.fileReport(p)
			if e != nil {
				return err
			}

			stats = append(stats, r...)
		}

		return err
	})
	return stats, nil
}

func (a analyzer) fileReport(path string) ([]stat, error) {
	stats := []stat{}
	if !strings.HasSuffix(path, ".go") {
		return stats, nil
	}

	if a.noTests && strings.HasSuffix(path, "_test.go") {
		return stats, nil
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, 0)
	if err != nil {
		return nil, err
	}

	for _, decl := range f.Decls {
		if fn, ok := decl.(*ast.FuncDecl); ok {
			stats = append(stats, stat{
				Pkg:        f.Name.Name,
				FuncName:   funcName(fn),
				Complexity: complexity(fn),
				Position:   fset.Position(fn.Pos()),
			})
		}
	}

	return stats, nil
}

type stat struct {
	Pkg        string
	FuncName   string
	Complexity int
	Position   token.Position
}

func funcName(fn *ast.FuncDecl) string {
	if fn.Recv != nil {
		if fn.Recv.NumFields() > 0 {
			typ := fn.Recv.List[0].Type
			return fmt.Sprintf("(%s).%s", recvString(typ), fn.Name)
		}
	}
	return fn.Name.Name
}

func recvString(recv ast.Expr) string {
	switch t := recv.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		return "*" + recvString(t.X)
	}
	return "BADRECV"
}

func complexity(fn *ast.FuncDecl) int {
	v := complexityVisitor{}
	ast.Walk(&v, fn)
	return v.Complexity
}

type complexityVisitor struct {
	Complexity int
}

func (v *complexityVisitor) Visit(n ast.Node) ast.Visitor {
	switch n := n.(type) {
	case *ast.FuncDecl, *ast.IfStmt, *ast.ForStmt, *ast.RangeStmt, *ast.CaseClause, *ast.CommClause:
		v.Complexity++
	case *ast.BinaryExpr:
		if n.Op == token.LAND || n.Op == token.LOR {
			v.Complexity++
		}
	}
	return v
}

func isDir(filename string) bool {
	fi, err := os.Stat(filename)
	return err == nil && fi.IsDir()
}
