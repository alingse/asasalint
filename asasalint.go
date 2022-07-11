package asasalint

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const BuiltinExclude = `Printf,Println,Fprintf,Fprintln,Fatal,Fatalf,Panic,Panicf,Panicln,Print,Printf,Println,Sprintf,Sprintln,Error,Errorf,Info,Infof,Warn,Warnf,Debug,Debugf`

type LinterSetting struct {
	Exclude          []string
	NoBuiltinExclude bool
	IgnoreInTest     bool
}

func NewAnalyzer(setting LinterSetting) *analysis.Analyzer {
	a := newAnalyzer(setting)
	return &analysis.Analyzer{
		Name:     "asasalint",
		Doc:      "check for pass []any as any in variadic func(...any)",
		Run:      a.run,
		Requires: []*analysis.Analyzer{inspect.Analyzer},
	}
}

type analyzer struct {
	excludes map[string]bool
	setting  LinterSetting
}

func newAnalyzer(setting LinterSetting) *analyzer {
	a := &analyzer{
		excludes: make(map[string]bool),
		setting:  setting,
	}

	if !a.setting.NoBuiltinExclude {
		for _, exclude := range strings.Split(BuiltinExclude, `,`) {
			a.excludes[exclude] = true
		}
	}

	for _, exclude := range a.setting.Exclude {
		if exclude != "" {
			a.excludes[exclude] = true
		}
	}

	return a
}

func (a *analyzer) run(pass *analysis.Pass) (interface{}, error) {
	inspectorInfo := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{(*ast.CallExpr)(nil)}

	inspectorInfo.Nodes(nodeFilter, a.AsCheckVisitor(pass))
	return nil, nil
}

func (a *analyzer) AsCheckVisitor(pass *analysis.Pass) func(n ast.Node, push bool) bool {
	return func(n ast.Node, push bool) (processed bool) {
		processed = true
		if a.setting.IgnoreInTest {
			pos := pass.Fset.Position(n.Pos())
			if strings.HasSuffix(pos.Filename, "_test.go") {
				return
			}
		}

		caller, ok := n.(*ast.CallExpr)
		if !ok {
			return
		}
		if caller.Ellipsis != token.NoPos {
			return
		}
		if len(caller.Args) == 0 {
			return
		}
		fnName := getFuncName(caller)
		if a.excludes[fnName] {
			return
		}

		fnType := pass.TypesInfo.TypeOf(caller.Fun)
		if !isSliceAnyVariadicFuncType(fnType) {
			return
		}

		fnSign := fnType.(*types.Signature)
		if len(caller.Args) != fnSign.Params().Len() {
			return
		}

		lastArg := caller.Args[len(caller.Args)-1]
		argType := pass.TypesInfo.TypeOf(lastArg)
		if !isSliceAnyType(argType) {
			return
		}
		node := lastArg

		d := analysis.Diagnostic{
			Pos: node.Pos(),
			End: node.End(),
			Message: fmt.Sprintf("pass []any as any to func %s %s",
				fnName, fnType.String()),
			Category: "asasalint",
		}
		pass.Report(d)
		return
	}
}

func getFuncName(caller *ast.CallExpr) string {
	if id, ok := caller.Fun.(*ast.Ident); ok {
		return id.Name
	}
	if s, ok := caller.Fun.(*ast.SelectorExpr); ok {
		return s.Sel.Name
	}
	return ""
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
