package analyzer

import (
	"flag"
	"go/ast"
	"go/token"
	"strings"

	"golang.org/x/tools/go/analysis"
)

var (
	flagSet flag.FlagSet
)

var maxComplexity int
var skipTests bool

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name:  "cyclop",
		Doc:   "calculates cyclomatic complexity",
		Run:   run,
		Flags: flagSet,
	}
}

func init() {
	flagSet.IntVar(&maxComplexity, "maxComplexity", 10, "max complexity the function can have")
	flagSet.BoolVar(&skipTests, "skipTests", false, "should the linter execute on test files as well")
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		ast.Inspect(f, func(node ast.Node) bool {
			f, ok := node.(*ast.FuncDecl)
			if !ok {
				// we check function by function
				return true
			}

			if skipTests && testFunc(f) {
				return true
			}

			comp := complexity(f)
			if comp > maxComplexity {
				pass.Reportf(node.Pos(), "calculated cyclomatic complexity for function %s is %d, max is %d", f.Name.Name, comp, maxComplexity)
			}

			return true
		})
	}
	return nil, nil
}

func testFunc(f *ast.FuncDecl) bool {
	return strings.HasPrefix(f.Name.Name, "Test")
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
