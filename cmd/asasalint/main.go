package main

import (
	"flag"
	"strings"

	"github.com/alingse/asasalint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var extraExclude string
	var include string
	var disableDefaultExclude bool
	var ignoreInTest bool
	flag.StringVar(&extraExclude,
		"e",
		"",
		"Extra Exclude Func Names, like: FuncA,FuncB,Func",
	)
	flag.StringVar(&extraExclude, "i", "",
		"Must Include Func Names, like: FuncA,FuncB,Func")
	flag.BoolVar(&disableDefaultExclude, "no-default-exclude", false,
		"disbale the default exclude func names: "+asasalint.DefaultExclude)
	flag.BoolVar(&ignoreInTest, "ignore-in-test", false,
		"ingore case in * _test.go")
	flag.Parse()

	setting := asasalint.LinterSetting{
		Exclude:               strings.Split(extraExclude, ","),
		Include:               strings.Split(include, ","),
		DisableDefaultExclude: true,
	}
	singlechecker.Main(asasalint.NewAnalyzer(setting))
}
