package main

import (
	"flag"
	"strings"

	"github.com/alingse/asasalint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var extraExclude string
	var noBuiltinExclude bool
	var ignoreInTest bool
	flag.StringVar(&extraExclude,
		"e",
		"",
		"extra exclude func names, like: FuncA,append,Append",
	)
	flag.BoolVar(&noBuiltinExclude, "no-builtin-exclude", false,
		"disbale the builtin exclude func names: "+asasalint.BuiltinExclude)
	flag.BoolVar(&ignoreInTest, "ignore-in-test", false,
		"ingore case in *_test.go")
	flag.Parse()

	setting := asasalint.LinterSetting{
		Exclude:          strings.Split(extraExclude, ","),
		NoBuiltinExclude: noBuiltinExclude,
		IgnoreInTest:     ignoreInTest,
	}
	singlechecker.Main(asasalint.NewAnalyzer(setting))
}
