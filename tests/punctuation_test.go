package tests

import (
	"testing"

	functions "go-reloaded/functions"
)

func TestPunctuation(t *testing.T) {
	tests := []struct {
		id       int
		input    string
		expected string
	}{
		{7, "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{8, "I was thinking ... You were right", "I was thinking... You were right"},
		{9, "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{10, "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{18, "This is ' tricky '", "This is 'tricky'"},
		{25, "Multiple ' quotes ' in ' one ' line", "Multiple 'quotes' in 'one' line"},
		{32, "... (up)", "..."}, // Requires pipeline processing to remove (up)
		{33, "word , (up)", "word, (up)"},
		{39, "word . . . word", "word... word"},
		{40, "word ,,,", "word,,,"},
		{41, "boom !!! ?", "boom!!!?"},
		{42, "' ' ' '", "'' ''"},
		{47, "a ' big ' apple", "a 'big' apple"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// 1. Tokenize
			tokens := functions.Tokenize(tt.input)

			// 2. Process Markers (Hex/Bin/Casing)
			// We include this because some Golden Tests (like #32) depend on markers being removed.
			tokens = functions.ProcessMarkers(tokens)
			tokens = functions.ProcessCasing(tokens)

			// 3. Process Punctuation (The SUT)
			tokens = functions.ProcessPunctuation(tokens)

			// 4. Compare by joining tokens to match the expected string (since Punctuation merges tokens)

			gotString := ""
			for i, tok := range tokens {
				if i > 0 {
					gotString += " "
				}
				gotString += tok
			}

			if gotString != tt.expected {
				t.Errorf("ID %d: \nInput:    %q\nGot:      %q\nExpected: %q", tt.id, tt.input, gotString, tt.expected)
			}
		})
	}
}
