package main

import (
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
)

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`analyze path "(.+)"`, analyze)
	suite.AddStep(`set top parameter to (\d+)`, setTopParameterTo)
	suite.AddStep(`it returns no error`, itReturnsNoError)
	suite.AddStep(`cyclomatic complexity of function ([\(\)\.a-zA-Z0-9]+) equals (\d+)`, cyclomaticComplexityOfFunctionEquals)
	suite.AddStep(`the size of the result should equal (\d+)`, theSizeOfResultShouldEqual)
	suite.Run()
}

type statsKey struct{}
type statsErrKey struct{}

func analyze(t gobdd.StepTest, ctx context.Context, filePath string) context.Context {
	c := NewCyclop()
	if top, err := ctx.GetInt("top"); err == nil && top > 0 {
		c = c.WithTopResults(top)
	}

	stats, err := c.AnalyzePaths([]string{filePath})
	ctx.Set(statsKey{}, stats)
	ctx.Set(statsErrKey{}, err)
	return ctx
}

func itReturnsNoError(t gobdd.StepTest, ctx context.Context) context.Context {
	res, _ := ctx.GetError(statsErrKey{})
	if res != nil {
		t.Errorf("expected no error but %+v received", res)
	}

	return ctx
}

func cyclomaticComplexityOfFunctionEquals(t gobdd.StepTest, ctx context.Context, f string, c int) context.Context {
	res, err := ctx.Get(statsKey{})
	if err != nil {
		t.Errorf("expected no error but %+v received", err)
		return ctx
	}

	stats := res.([]stat)

	for _, s := range stats {
		if s.FuncName == f {
			if s.Complexity != c {
				t.Errorf("expected complexity %d bug %d given", c, s.Complexity)
			}

			return ctx
		}
	}

	t.Errorf("could not find statistics for function %s", f)
	return ctx
}

func setTopParameterTo(t gobdd.StepTest, ctx context.Context, top int) context.Context {
	ctx.Set("top", top)
	return ctx
}

func theSizeOfResultShouldEqual(t gobdd.StepTest, ctx context.Context, s int) context.Context {
	res, err := ctx.Get(statsKey{})
	if err != nil {
		t.Errorf("expected no error but %+v received", res)
	}

	stats := res.([]stat)

	if len(stats) != s {
		t.Errorf("expected %d elements but %d received", s, len(stats))
	}

	return ctx
}
