package main

import "testing"

var cases = []struct {
	input    string
	expected []string
}{
	{
		input:    "  hello  world  ",
		expected: []string{"hello", "world"},
	},
	{
		input:    "",
		expected: []string{},
	},
	{
		input:    "hEllo wOrLD",
		expected: []string{"hello", "world"},
	},
	// add more cases here
}

func TestCleanInput(t *testing.T) {
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Error: for input: %v. Length actual: %v vs len expected: %v", c.input, actual, c.expected)
			continue
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Error: word: %v vs expectedWord: %v", word, expectedWord)
			}
		}
	}
}
