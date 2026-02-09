package tests

import (
	"reflect"
	"testing"

	functions "go-reloaded/functions"
)

func TestCasingMarkers(t *testing.T) {
	tests := []struct {
		id       int
		input    string
		expected string
	}{
		{3, "Ready, set, go (up) !", "Ready, set, GO !"},
		{4, "I should stop SHOUTING (low)", "I should stop shouting"},
		{5, "Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{6, "This is so exciting (up, 2)", "This is SO EXCITING"},
		{14, "Please make this loud (up)", "Please make this LOUD"},
		{15, "This is a MAJOR (low) mistake", "This is a major mistake"},
		{16, "my name is john doe (cap, 2)", "my name is John Doe"},
		{20, "(up) starts with marker", "starts with marker"},
		{21, "Hello (up, 5) world", "HELLO world"},
		{23, "User (cap, 0) test", "User test"},
		{26, "Nested (hex) (up)", "Nested"},
		{28, "1e (hex) (up) is thirty", "30 is thirty"},
		{29, "it (cap) (up) worked", "IT worked"},
		{34, "one (up) (low) (cap)", "One"},
		{38, "hello world (up, 2) (low, 1)", "HELLO world"},
		{44, "42 (up)", "42"},
		{45, "word (nonsense)", "word (nonsense)"},
		{46, "(up, -1)", ""},
		{49, "hello (up, 100) world", "HELLO world"},
		{50, "sample (cap) (cap) (cap)", "Sample"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// 1. Tokenize
			tokens := functions.Tokenize(tt.input)

			// 2. Process Markers (Hex/Bin) first.
			// Some tests (e.g. #28) rely on hex conversion happening before casing.
			tokens = functions.ProcessMarkers(tokens)

			// 3. Process Casing (The SUT)
			tokens = functions.ProcessCasing(tokens)

			// Filter empty strings (placeholders) for comparison
			cleanTokens := []string{}
			for _, t := range tokens {
				if t != "" {
					cleanTokens = append(cleanTokens, t)
				}
			}
			tokens = cleanTokens

			// 4. Compare
			expectedTokens := functions.Tokenize(tt.expected)
			if !reflect.DeepEqual(tokens, expectedTokens) {
				t.Errorf("ID %d: Input %q\nGot:  %v\nWant: %v", tt.id, tt.input, tokens, expectedTokens)
			}
		})
	}
}
