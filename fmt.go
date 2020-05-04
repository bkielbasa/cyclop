package main

import "fmt"

func display(stats []stat) {
	if *shortAvg == true {
		fmt.Printf("%.3g\n", average(stats))
		return
	}
	if *top == 0 {
		displayStats(stats)
	} else {
		displayTopStats(stats, *top)
	}
}

func displayStats(stats []stat) {
	for _, s := range stats {
		fmt.Printf("%d %s.%s\n", s.Complexity, s.Pkg, s.FuncName)
	}
}

func displayTopStats(stats []stat, top int) {
	for i, s := range stats {
		if i == top {
			return
		}

		fmt.Printf("%d %s.%s\n", s.Complexity, s.Pkg, s.FuncName)
	}
}
