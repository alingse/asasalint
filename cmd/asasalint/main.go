package main

import (
	"flag"
	"strings"

	"github.com/alingse/asasalint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var extraExclude string
	var noDefaultExclude bool
	var ignoreInTest bool
	flag.StringVar(&extraExclude,
		"e",
		"",
		"Extra Exclude Func Names, like: FuncA,append,Append",
	)
	flag.BoolVar(&noDefaultExclude, "no-default-exclude", false,
		"disbale the default exclude func names: "+asasalint.DefaultExclude)
	flag.BoolVar(&ignoreInTest, "ignore-in-test", false,
		"ingore case in *_test.go")
	flag.Parse()

	setting := asasalint.LinterSetting{
		Exclude:          strings.Split(extraExclude, ","),
		NoDefaultExclude: noDefaultExclude,
		IgnoreInTest:     ignoreInTest,
	}
	singlechecker.Main(asasalint.NewAnalyzer(setting))
}
