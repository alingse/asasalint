package asasalint

import (
	"fmt"
	"go/ast"
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

	fn := s.Pass.TypesInfo.TypeOf(caller.Fun)
	fnSign := fn.(*types.Signature)
	if !fnSign.Variadic() {
		return true
	}

	params := fnSign.Params()
	param := params.At(params.Len() - 1)
	varParamType := param.Type().(*types.Slice)
	paramType, ok := varParamType.Elem().(*types.Interface)
	if !ok || paramType.NumMethods() != 0 {
		return true
	}
	// assert the last param is []any
	fmt.Printf("paramType %#v\n", paramType)
	fmt.Println("param", param)
	return false
}
