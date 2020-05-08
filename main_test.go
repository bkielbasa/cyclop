package main

import (
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
)

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`analyze path "(.+)"`, analyzeFile)
	suite.AddStep(`it returns no error`, itReturnsNoError)
	suite.AddStep(`cyclomatic complexity of function ([a-zA-Z0-9]+) equals (\d+)`, cyclomaticComplexityOfFunctionEquals)
	suite.Run()
}

type statsKey struct{}
type statsErrKey struct{}

func analyzeFile(t gobdd.StepTest, ctx context.Context, filePath string) context.Context {
	a := analyzer{}
	stats, err := a.analyze([]string{filePath})
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
