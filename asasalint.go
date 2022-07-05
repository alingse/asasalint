package asasalint

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"
	"sync"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var defaultExcludes []string

func init() {
	defaultExcludes = strings.Split(
		`Printf,Println,Errorf,Fprintf,Fprintln,Fatal,Fatalf,Panic,Panicf,Panicln,Print,Printf,Println,Sprintf,Sprintln`+
			`Error,Errorf,Info,Infof,Warn,Warnf,Debug,Debugf`, `,`)
}

func NewAnalyzer(excludes []string, include []string) *analysis.Analyzer {
	a := &analyzer{
		excludes: make(map[string]bool),
	}
	for _, exclude := range defaultExcludes {
		a.excludes[exclude] = true
	}
	for _, exclude := range excludes {
		a.excludes[exclude] = true
	}

	for _, include := range include {
		a.excludes[include] = false
	}

	return &analysis.Analyzer{
		Name:     "asasalint",
		Doc:      "check for pass []any as any in func(...any)",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

type analyzer struct {
	excludes map[string]bool
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	inspectorInfo := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.CallExpr)(nil),
	}
	search := &Searcher{Pass: pass}
	inspectorInfo.Nodes(nodeFilter, search.CheckAndReport)
	search.wg.Wait()
	return nil, nil
}

type Searcher struct {
	Pass *analysis.Pass
	wg   sync.WaitGroup
}

func (s *Searcher) CheckAndReport(n ast.Node, push bool) bool {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		caller, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}
		if caller.Ellipsis != token.NoPos {
			return
		}

		fnType := s.Pass.TypesInfo.TypeOf(caller.Fun)
		// fmt.Println("process this func --> ", fnType.String())
		if !isSliceAnyVariadicFuncType(fnType) {
			return
		}
		if len(caller.Args) == 0 {
			return
		}

		lastArg := caller.Args[len(caller.Args)-1]
		argType := s.Pass.TypesInfo.TypeOf(lastArg)
		if !isSliceAnyType(argType) {
			return
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
	}()

	return true
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
