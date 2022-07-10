package asasalint

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"
)

const filename = "<src>"

func makePkg(src string) (*types.Package, error) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, src, parser.DeclarationErrors)
	if err != nil {
		return nil, err
	}
	// use the package name as package path
	conf := types.Config{Importer: importer.Default()}
	return conf.Check(file.Name.Name, fset, []*ast.File{file}, nil)
}

func TestIsSliceAnyType(t *testing.T) {
	p, _ := makePkg("package p; var A []any\nvar B any = []any{1, 2, 3}\nvar C []int")
	aT := p.Scope().Lookup("A").Type()
	if !isSliceAnyType(aT) {
		t.Errorf("isSliceAnyType(%v) = false, want true", aT)
	}
	bT := p.Scope().Lookup("B").Type()
	if isSliceAnyType(bT) {
		t.Errorf("isSliceAnyType(%v) = true, want false", bT)
	}
	cT := p.Scope().Lookup("C").Type()
	if isSliceAnyType(cT) {
		t.Errorf("isSliceAnyType(%v) = true, want false", cT)
	}
}

func TestIsSliceAnyVariadicFuncType(t *testing.T) {
	p, _ := makePkg("package p; func hello(a int, b ...any) {}\nfunc hello2(a int, b int) {}")
	ft := p.Scope().Lookup("hello").Type()
	if !isSliceAnyVariadicFuncType(ft) {
		t.Errorf("isSliceAnyVariadicFuncType(%v) = false, want true", ft)
	}

	ft2 := p.Scope().Lookup("hello2").Type()
	if isSliceAnyVariadicFuncType(ft2) {
		t.Errorf("isSliceAnyVariadicFuncType(%v) = true, want false", ft2)
	}
}

func TestGetFuncName(t *testing.T) {
	src := `package p
func hello(a int, b ...any) {}
func hello2(a int, b int) { hello(a, b)}
type A struct {}
func (a *A) hello(a int, b ...any) {
	hello(a, b...)
}
func (a *A) hello2(a int, b int) {
	a.hello(a, b)
}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "src.go", src, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.CallExpr:
			s := getFuncName(x)
			if s != "hello" {
				t.Errorf("getFuncName(%#v) = %v, want hello", x.Fun, s)
			}
		}
		return true
	})
}

func TestNewAnalyzer(t *testing.T) {
	_ = NewAnalyzer(LinterSetting{})
	_ = NewAnalyzer(LinterSetting{
		Exclude:               []string{"hello"},
		Include:               []string{"hello", "world"},
		DisableDefaultExclude: true,
		IgnoreInTest:          true,
	})
}
