package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello  World  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Got %s, expected %s", word, expectedWord)
			}
		}
	}
}
