package main

import (
	"github.com/alingse/asasalint"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(asasalint.Analyzer)
}
