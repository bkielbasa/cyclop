package p

import "testing"

func TestIfTestFunctionsArentSkipped(t *testing.T) { // want "calculated cyclomatic complexity for function"
	i := 1
	if i > 2 {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	} else {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	}

	if i > 2 {
	}
}

func noComplexity() {}

func highComplexity() { // want "calculated cyclomatic complexity for function"
	i := 1
	if i > 2 {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	} else {
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
		if i > 2 {
		}
	}

	if i > 2 {
	}
}
