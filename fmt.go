package main

import "fmt"

func display(stats []stat) {
	if *shortAvg == true {
		fmt.Printf("%.3g\n", average(stats))
		return
	}

	displayStats(stats)
}

func displayStats(stats []stat) {
	for _, s := range stats {
		fmt.Printf("%d %s.%s\n", s.Complexity, s.Pkg, s.FuncName)
	}
}
