package tests

import (
	"reflect"
	"testing"

	functions "go-reloaded/functions"
)

func TestHexBinMarkers(t *testing.T) {
	tests := []struct {
		id       int
		input    string
		expected string
	}{
		{1, "1E (hex) files were added", "30 files were added"},
		{2, "It has been 10 (bin) years", "It has been 2 years"},
		{12, "I bought 1010 (hex) apples", "I bought 4112 apples"},
		{13, "I have known him for 10100 (bin) years", "I have known him for 20 years"},
		{22, "(hex) 1E", "(hex) 1E"}, // Ignored (start of line)
		{31, "He said: ' 10 (bin) '", "He said: ' 2 '"},
		{37, "10 (bin) (hex)", "2"},                  // Chain: 10(bin)->2, 2(hex)->2
		{43, "(hex) (bin) (up)", "(hex) (bin) (up)"}, // Orphans
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// 1. Tokenize input (using the already implemented Tokenizer)
			tokens := functions.Tokenize(tt.input)

			// 2. Process Markers (The function we are testing)
			gotTokens := functions.ProcessMarkers(tokens)

			// 3. Compare with expected output (Tokenized)
			// We tokenize the expected string to compare slice vs slice
			expectedTokens := functions.Tokenize(tt.expected)

			if !reflect.DeepEqual(gotTokens, expectedTokens) {
				t.Errorf("ID %d: ProcessMarkers(%q) = %v, want %v", tt.id, tt.input, gotTokens, expectedTokens)
			}
		})
	}
}
