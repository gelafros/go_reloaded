package tests

import (
	"testing"

	functions "go-reloaded/functions"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		id       int
		input    string
		expected string
	}{
		// 1. Audit Cases
		{1, "1E (hex) files were added", "30 files were added"},
		{2, "It has been 10 (bin) years", "It has been 2 years"},
		{3, "Ready, set, go (up) !", "Ready, set, GO!"},
		{4, "I should stop SHOUTING (low)", "I should stop shouting"},
		{5, "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{6, "This is so exciting (up, 2)", "This is SO EXCITING"},
		{7, "I was sitting over there ,and then BAMM !!", "I was sitting over there, and then BAMM!!"},
		{8, "I was thinking ... You were right", "I was thinking... You were right"},
		{9, "I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{10, "As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
		{11, "There it was. A amazing rock!", "There it was. An amazing rock!"},

		// 5. The Gauntlet (Complex Interactions)
		{34, "one (up) (low) (cap)", "One"},
		{35, "a (up) apple", "AN apple"},
		{36, "a (low) apple", "an apple"},
		{37, "10 (bin) (hex)", "2"},
		{38, "hello world (up, 2) (low, 1)", "HELLO world"},
		{39, "word . . . word", "word... word"},
		{40, "word ,,,", "word,,,"},
		{41, "boom !!! ?", "boom!!!?"},
		{42, "' ' ' '", "'' ''"},
		{43, "(hex) (bin) (up)", ""},
		{44, "42 (up)", "42"},
		{45, "word (nonsense)", "word (nonsense)"},
		{46, "(up, -1)", ""},
		{47, "a ' big ' apple", "a 'big' apple"},
		{48, "tok&en (up)", "TOK&EN"},
		{49, "hello (up, 100) world", "HELLO world"},
		{50, "sample (cap) (cap) (cap)", "Sample"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := functions.Process(tt.input)
			if got != tt.expected {
				t.Errorf("ID %d:\nInput:    %q\nGot:      %q\nExpected: %q", tt.id, tt.input, got, tt.expected)
			}
		})
	}
}
