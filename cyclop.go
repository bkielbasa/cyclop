package main

import (
	"fmt"
	"regexp"
	"strings"
)

func NewCyclop() Cyclop {
	return Cyclop{}
}

type Cyclop struct {
	analyzer  analyzer
	skipTests bool
	top       int
}

func (c Cyclop) AnalyzePaths(paths []string, ignorePattern string) ([]stat, error) {
	stats, err := c.analyzer.analyze(paths)

	if c.skipTests {
		stats = c.filterOutTests(stats)
	}

	if ignorePattern != "" {
		reg, err := regexp.Compile(ignorePattern)
		if err != nil {
			return nil, fmt.Errorf("cannot parse regexp %s: %w", ignorePattern, err)
		}
		stats = c.filterOutFilePattern(stats, reg)
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

func (c Cyclop) filterOutFilePattern(stats []stat, reg *regexp.Regexp) []stat {
	res := []stat{}

	for _, s := range stats {
		if reg.MatchString(s.Position.Filename) {
			continue
		}

		res = append(res, s)
	}

	return res
}
func (c Cyclop) filterOutTests(stats []stat) []stat {
	res := []stat{}

	for _, s := range stats {
		if strings.HasPrefix(s.FuncName, "Test") {
			continue
		}

		res = append(res, s)
	}

	return res
}

func maxTop(s []stat, max int) []stat {
	if max > len(s) {
		max = len(s)
	}
	return s[:max]
}
