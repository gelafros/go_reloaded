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
		{20, "(up) starts with marker", "(up) starts with marker"},
		{21, "Hello (up, 5) world", "HELLO world"},
		{23, "User (cap, 0) test", "User test"},
		{26, "Nested (hex) (up)", "Nested (hex) (up)"},
		{28, "1e (hex) (up) is thirty", "30 is thirty"},
		{29, "it (cap) (up) worked", "IT worked"},
		{34, "one (up) (low) (cap)", "One"},
		{38, "hello world (up, 2) (low, 1)", "HELLO world"},
		{44, "42 (up)", "42"},
		{45, "word (nonsense)", "word (nonsense)"},
		{46, "(up, -1)", "(up, -1)"},
		{49, "hello (up, 100) world", "HELLO world"},
		{50, "sample (cap) (cap) (cap)", "Sample"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			// 1. Tokenize
			tokens := functions.Tokenize(tt.input)

			// 2. Process Markers (Hex/Bin first as per pipeline, but here we test Casing specifically)
			// However, test #28 "1e (hex) (up)" requires Hex to run first to become "30", then "30" (up) -> "30".
			// If we only run Casing, "1e" (hex) (up) -> "1e" (hex) (up) (since (hex) isn't a casing marker).
			// Wait, the test expects "30 is thirty". This implies the full pipeline or at least Hex+Casing.
			// For unit testing Casing, we should assume Hex is already done OR we are testing the integration of markers.
			// Let's assume we are testing `ProcessCasing`.
			// BUT, for test #28, if we pass "1e (hex) (up)", `ProcessCasing` will see "(hex)" as the word to uppercase?
			// No, `(hex)` is a marker.
			// To strictly test Casing in isolation, we might need to pre-process Hex for that specific test,
			// OR we accept that `ProcessCasing` handles the casing logic and we might need to chain them in the test
			// if the test input depends on it.
			// Actually, `ProcessMarkers` in TASK-04 handled Hex/Bin.
			// TASK-05 is `functions/casing.go`.
			// Let's chain them in the test to ensure "Golden Test" compliance.

			tokens = functions.ProcessMarkers(tokens) // Hex/Bin
			tokens = functions.ProcessCasing(tokens)  // Casing (The SUT)

			// 3. Compare
			expectedTokens := functions.Tokenize(tt.expected)
			if !reflect.DeepEqual(tokens, expectedTokens) {
				t.Errorf("ID %d: Input %q\nGot:  %v\nWant: %v", tt.id, tt.input, tokens, expectedTokens)
			}
		})
	}
}
