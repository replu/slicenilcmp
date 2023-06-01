package main

import (
	"github.com/replu/slicenilcmp"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(slicenilcmp.Analyzer) }
