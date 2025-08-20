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
			input:    "  Hello  World   ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " has, comma",
			expected: []string{"has,", "comma"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Input: %s | Expected length %d, got length %d", c.input, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Input: %s | index %d | expected '%s', actual '%s'", c.input, i, expectedWord, word)
			}
		}
	}
}
