package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "banana land",
			expected: []string{"banana", "land"},
		},
		{
			input:    "n0 fear shakespeare",
			expected: []string{"n0", "fear", "shakespeare"},
		},
		{
			input:    "<F19>65 dskj *(_)",
			expected: []string{"<f19>65", "dskj", "*(_)"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Fatalf("Error: slice lengths do not match!\nActual: %d, Expected: %d", len(actual), len(c.expected))
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error: processed words do not match!\n'%s' compared to '%s'", word, expectedWord)
			}
		}
	}
}
