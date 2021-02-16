package main

import (
	"github.com/bkielbasa/cyclop/v2/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.NewAnalyzer())
}
