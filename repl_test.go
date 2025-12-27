package main

import (
	"testing"
	"reflect"
)

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
			input:    "dark  souls",
			expected: []string{"dark", "souls"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "a        a",
			expected: []string{"a","a"},
		},
	// add more cases here
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if !reflect.DeepEqual(len(c.expected), len(actual)) {
			t.Errorf("expected: %v, got: %v", len(c.expected), len(actual))
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if !reflect.DeepEqual(word, expectedWord) {
         		t.Errorf("expected: %v, got: %v", word, expectedWord)
				return
    		}
		}
	}
}