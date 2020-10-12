package main

import (
	"bytes"
	"testing"

	"github.com/go-bdd/gobdd"
	"github.com/go-bdd/gobdd/context"
	"github.com/stretchr/testify/assert"
)

func TestScenarios(t *testing.T) {
	suite := gobdd.NewSuite(t)
	suite.AddStep(`analyze path "(.+)"`, analyze)
	suite.AddStep(`set top parameter to (\d+)`, setTopParameterTo)
	suite.AddStep(`set skipping tests`, setSkippingTests)
	suite.AddStep(`it returns no error`, itReturnsNoError)
	suite.AddStep(`cyclomatic complexity of function ([\(\)\.a-zA-Z0-9]+) equals (\d+)`, cyclomaticComplexityOfFunctionEquals)
	suite.AddStep(`the size of the result should equal (\d+)`, theSizeOfResultShouldEqual)
	suite.Run()
}

func TestRun(t *testing.T) {
	var buf []byte
	b := bytes.NewBuffer(buf)
	err := run(b, []string{"internal/simple.go"})

	assert.NoError(t, err)
	assert.Equal(t, "3 internal.Or\n3 internal.And\n2 internal.OneIf\n1 internal.(*S).PointerFunction\n1 internal.(S).AFunction\n1 internal.NoComplexity\n", b.String())
}

func TestRunWitTotal(t *testing.T) {
	var buf []byte
	b := bytes.NewBuffer(buf)
	err := run(b, []string{"-total=true", "internal/simple.go"})

	assert.NoError(t, err)
	assert.Equal(t, "Total: 11\n", b.String())
}

type statsKey struct{}
type statsErrKey struct{}

func analyze(t gobdd.StepTest, ctx context.Context, filePath string) context.Context {
	c := NewCyclop()
	if top, err := ctx.GetInt("top"); err == nil && top > 0 {
		c = c.WithTopResults(top)
	}

	if skip, err := ctx.GetBool("skip-tests"); err == nil && skip == true {
		c = c.WithNoTests()
	}

	stats, err := c.AnalyzePaths([]string{filePath}, "")
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

func setSkippingTests(t gobdd.StepTest, ctx context.Context) context.Context {
	ctx.Set("skip-tests", true)
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
