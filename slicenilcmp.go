package slicenilcmp

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "slicenil is ..."

var Analyzer = &analysis.Analyzer{
	Name: "slicenil",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.BinaryExpr)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		be, ok := n.(*ast.BinaryExpr)
		if !ok {
			return
		}

		if !(be.Op == token.EQL || be.Op == token.NEQ) {
			return
		}

		if !isSlice(pass, be.X, be.Y) {
			return
		}

		if !isNil(pass, be.X, be.Y) {
			return
		}

		pass.Reportf(n.Pos(), "suggestion: use len func for empty check")
	})

	return nil, nil
}

func isSlice(pass *analysis.Pass, x, y ast.Expr) bool {
	xType := pass.TypesInfo.TypeOf(x)
	if xType == nil {
		return false
	}
	if _, ok := xType.(*types.Slice); ok {
		return true
	}

	yType := pass.TypesInfo.TypeOf(y)
	if yType == nil {
		return false
	}
	if _, ok := yType.(*types.Slice); ok {
		return true
	}
	return false
}

func isNil(pass *analysis.Pass, x, y ast.Expr) bool {
	nilType := types.Typ[types.UntypedNil]

	xType := pass.TypesInfo.TypeOf(x)
	if xType == nil {
		return false
	}
	if types.Identical(xType, nilType) {
		return true
	}

	yType := pass.TypesInfo.TypeOf(y)
	if yType == nil {
		return false
	}
	if types.Identical(yType, nilType) {
		return true
	}

	return false
}
