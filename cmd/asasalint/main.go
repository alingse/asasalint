package main

import (
	"flag"
	"log"
	"strings"

	"github.com/alingse/asasalint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var extraExclude string
	var noBuiltinExclusions bool
	var ignoreTest bool
	flag.StringVar(&extraExclude, "e", "",
		"Extra exclusions func names. It can be regular expressions. ex: FuncA,(A|a)ppend")
	flag.BoolVar(&noBuiltinExclusions, "no-builtin-exclude", false,
		"Disable the builtin exclusions func names: "+asasalint.BuiltinExclusions)
	flag.BoolVar(&ignoreTest, "ignore-test", false,
		"Ignore test files (*_test.go)")
	flag.Parse()

	setting := asasalint.LinterSetting{
		Exclude:             strings.Split(extraExclude, ","),
		NoBuiltinExclusions: noBuiltinExclusions,
		IgnoreTest:          ignoreTest,
	}

	analyzer, err := asasalint.NewAnalyzer(setting)
	if err != nil {
		log.Fatal(err)
	}

	singlechecker.Main(analyzer)
}
