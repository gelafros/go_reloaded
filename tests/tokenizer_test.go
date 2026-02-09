package tests

import (
	"reflect"
	"testing"

	functions "go-reloaded/functions"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		id       int
		input    string
		expected []string
	}{
		{
			id:    6,
			input: "This is so exciting (up, 2)",
			// Strict rule: comma is padded. "(up, 2)" -> "(up", ",", "2)"
			expected: []string{"This", "is", "so", "exciting", "(up", ",", "2)"},
		},
		{
			id:       7,
			input:    "I was sitting over there ,and then BAMM !!",
			expected: []string{"I", "was", "sitting", "over", "there", ",", "and", "then", "BAMM", "!", "!"},
		},
		{
			id:       8,
			input:    "I was thinking ... You were right",
			expected: []string{"I", "was", "thinking", "...", "You", "were", "right"},
		},
		{
			id:       39,
			input:    "word . . . word",
			expected: []string{"word", ".", ".", ".", "word"},
		},
		{
			id:       40,
			input:    "word ,,,",
			expected: []string{"word", ",", ",", ","},
		},
		{
			id:       41,
			input:    "boom !!! ?",
			expected: []string{"boom", "!", "!", "!", "?"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := functions.Tokenize(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("ID %d: Tokenize(%q) = %v, want %v", tt.id, tt.input, got, tt.expected)
			}
		})
	}
}
