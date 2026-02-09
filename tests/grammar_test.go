package tests

import (
	"testing"

	functions "go-reloaded/functions"
)

func TestGrammar(t *testing.T) {
	tests := []struct {
		id       int
		input    string
		expected string
	}{
		{11, "There it was. A amazing rock!", "There it was. An amazing rock!"},
		{19, "It has a high value", "It has an high value"},
		{30, "A (up) amazing result", "AN amazing result"},
		{35, "a (up) apple", "AN apple"},
		{36, "a (low) apple", "an apple"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// 1. Tokenize
			tokens := functions.Tokenize(tt.input)

			// 2. Pre-processing (Markers & Punctuation)
			tokens = functions.ProcessMarkers(tokens)
			// Grammar must run before Casing to handle "a (up)" -> "AN" correctly.
			// It skips the (up) marker to check the next word, converts a->an,
			// and then Casing converts an->AN.
			tokens = functions.ProcessGrammar(tokens)
			tokens = functions.ProcessCasing(tokens)
			tokens = functions.ProcessPunctuation(tokens)

			// 4. Compare (Join tokens to match expected string)
			gotString := ""
			for i, tok := range tokens {
				if i > 0 {
					gotString += " "
				}
				gotString += tok
			}

			// Note: Simple join works because Punctuation step handles spacing logic (merging)
			// and Grammar step just changes words.

			if gotString != tt.expected {
				t.Errorf("ID %d: \nInput:    %q\nGot:      %q\nExpected: %q", tt.id, tt.input, gotString, tt.expected)
			}
		})
	}
}
