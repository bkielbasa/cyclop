package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	noTests  = flag.Bool("no-tests", true, "should ignore test files")
	top      = flag.Int("top", 0, "displays top N functions with the biggest complexity")
	failFrom = flag.Int("fail-from", 10, "returns the program with non-zero result if find at least one function with complexity higher than N")
	shortAvg = flag.Bool("short-avg", false, "displays only average complexity in short format")
)

func main() {
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}

	a := analyzer{
		noTests: *noTests,
		top:     *top,
	}

	stats, err := a.analyze(args)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	sortStats(stats)
	if len(stats) != 0 {
		display(stats)
	}

	if isFailed(stats) {
		os.Exit(1)
	}
}

func isFailed(stats []stat) bool {
	for _, s := range stats {
		if s.Complexity >= *failFrom {
			return true
		}
	}

	return false
}

func average(stats []stat) float64 {
	total := 0
	for _, s := range stats {
		total += s.Complexity
	}
	return float64(total) / float64(len(stats))
}
