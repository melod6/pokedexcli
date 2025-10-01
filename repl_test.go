package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input 	 string
		expected []string
	}{
		{
			input: "  hello world  ",
			expected: []string{"hello", "world"},
		},
		// more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c, input)
		// Check length of expected slice against actual slice
		// use t.Errorf to print error message and fail test if it doesn't match
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// check each word in the slice
			// if they don't match, use t.Errorf to print error message and fail the test
		}
	}
}
