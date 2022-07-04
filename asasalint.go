package asasalint

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "asasalint",
	Doc:  "check for pass []any as any in func(...any)",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	return nil, nil
}
