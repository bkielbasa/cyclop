package main

import "github.com/stretchr/testify/assert"
import "github.com/stretchr/testify/require"
import "testing"

func TestAnalize(t *testing.T) {
	a := analyzer{}
	stats, err := a.start([]string{"internal/simple.go"})
	expectedResults := map[string]int{
		"NoComplexity": 1,
		"OneIf":        2,
	}

	require.NoError(t, err)
	require.Equal(t, len(expectedResults), len(stats))

	for _, s := range stats {
		expected := expectedResults[s.FuncName]
		assert.Equal(t, expected, s.Complexity)
	}
}
