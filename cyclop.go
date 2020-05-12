package main

import "strings"

func NewCyclop() Cyclop {
	return Cyclop{}
}

type Cyclop struct {
	analyzer  analyzer
	skipTests bool
	top       int
}

func (c Cyclop) AnalyzePaths(paths []string) ([]stat, error) {
	stats, err := c.analyzer.analyze(paths)

	if c.skipTests {
		stats = c.filterOutTests(stats)
	}

	sortStats(stats)

	if c.top > 0 {
		stats = maxTop(stats, c.top)
	}

	return stats, err
}

func (c Cyclop) WithNoTests() Cyclop {
	c.skipTests = true
	return c
}

func (c Cyclop) WithTopResults(top int) Cyclop {
	c.top = top
	return c
}

func (c Cyclop) filterOutTests(stats []stat) []stat {
	res := []stat{}

	for _, s := range stats {
		if strings.HasSuffix(s.FuncName, "Test") {
			continue
		}

		res = append(res, s)
	}

	return res
}

func maxTop(s []stat, max int) []stat {
	return s[:max]
}
