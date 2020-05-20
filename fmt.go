package main

import (
	"fmt"
	"io"
)

func display(w io.Writer, stats []stat) {
	if *shortAvg == true {
		fmt.Printf("%.3g\n", average(stats))
		return
	}

	displayStats(w, stats)
}

func displayStats(w io.Writer, stats []stat) {
	for _, s := range stats {
		_, _ = fmt.Fprint(w, fmt.Sprintf("%d %s.%s\n", s.Complexity, s.Pkg, s.FuncName))
	}
}

func displayTotal(w io.Writer, stats []stat) {
	t := 0

	for i := range stats {
		t += stats[i].Complexity
	}

	if *shortAvg {
		_, _ = fmt.Fprintf(w, fmt.Sprintf("%d\n", t))
	} else {
		_, _ = fmt.Fprintf(w, fmt.Sprintf("Total: %d\n", t))
	}
}