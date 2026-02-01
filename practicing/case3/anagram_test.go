package main

import "testing"

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		word1    string
		word2    string
		expected bool
	}{
		{"uva", "avu", true},
		{"abc", "cba", true},
		{"uva", "abc", false},
		{"UVA", "avu", false},
	}

	for _, testCase := range tests {
		result := IsAnagram(testCase.word1, testCase.word2)

		if result != testCase.expected {
			t.Errorf(
				"IsAnagram(%q, %q) = %v; want %v",
				testCase.word1,
				testCase.word2,
				result,
				testCase.expected,
			)
		}
	}
}
