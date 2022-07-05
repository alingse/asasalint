package main

import (
	"flag"
	"strings"

	"github.com/alingse/asasalint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	var extraExclude string
	flag.StringVar(&extraExclude,
		"e",
		"",
		"Extra Exclude Func Names, like: FuncA,FuncB,Func",
	)
	flag.Parse()
	excludes := strings.Split(extraExclude, ",")
	singlechecker.Main(asasalint.NewAnalyzer(excludes, nil))
}
