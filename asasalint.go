package asasalint

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name:             "asasalint",
	Doc:              "check for pass []any as any in func(...any)",
	Run:              run,
	RunDespiteErrors: true,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspectorInfo := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	search := &Searcher{
		Pass: pass,
	}

	inspectorInfo.WithStack(nodeFilter, search.CheckAndReport)
	return nil, nil
}

type Searcher struct {
	Pass *analysis.Pass
}

func (s *Searcher) CheckAndReport(n ast.Node, push bool, stack []ast.Node) bool {
	caller, ok := n.(*ast.CallExpr)
	if !ok {
		return true
	}
	if caller.Ellipsis != token.NoPos {
		return true
	}

	fnType := s.Pass.TypesInfo.TypeOf(caller.Fun)
	if !isSliceAnyVariadicFuncType(fnType) {
		return true
	}

	lastArg := caller.Args[len(caller.Args)-1]
	argType := s.Pass.TypesInfo.TypeOf(lastArg)
	if !isSliceAnyType(argType) {
		return true
	}
	node := lastArg
	// report a diagnostic
	d := analysis.Diagnostic{
		Pos:      node.Pos(),
		End:      node.End(),
		Message:  fmt.Sprintf("pass []any as any to %s", fnType.String()),
		Category: "asasalint",
	}
	s.Pass.Report(d)
	return false
}

func isSliceAnyVariadicFuncType(typ types.Type) (r bool) {
	fnSign, ok := typ.(*types.Signature)
	if !ok || !fnSign.Variadic() {
		return false
	}

	params := fnSign.Params()
	lastParam := params.At(params.Len() - 1)
	return isSliceAnyType(lastParam.Type())
}

func isSliceAnyType(typ types.Type) (r bool) {
	sliceType, ok := typ.(*types.Slice)
	if !ok {
		return
	}
	elemType, ok := sliceType.Elem().(*types.Interface)
	if !ok {
		return
	}
	return elemType.NumMethods() == 0
}
