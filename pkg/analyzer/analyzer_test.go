package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	skipTests = false

	testdata := filepath.Join(filepath.Dir(filepath.Dir(path)), "testdata")
	analysistest.Run(t, testdata, NewAnalyzer(), "p")
}

func TestIfTestFunctionsAreSkipped(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	skipTests = true

	testdata := filepath.Join(filepath.Dir(filepath.Dir(path)), "testdata")
	analysistest.Run(t, testdata, NewAnalyzer(), "tests")
}

func TestAverageComplexityOfAPackage(t *testing.T) {
	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	skipTests = false
	packageAverage = 5

	testdata := filepath.Join(filepath.Dir(filepath.Dir(path)), "testdata")
	analysistest.Run(t, testdata, NewAnalyzer(), "avg")
}
