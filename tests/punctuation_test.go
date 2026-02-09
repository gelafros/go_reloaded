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
		{32, "... (up)", "..."}, // In isolation, Punctuation sees ["...", "(up)"] -> ["...", "(up)"] if (up) is ignored?
		// Actually, for unit testing Punctuation, we should test the spacing logic.
		// If we pass raw string "... (up)", Tokenizer gives ["...", "(up)"].
		// ProcessPunctuation should return ["...", "(up)"] (no change to spacing).
		// But expected string is "...". This implies (up) is gone.
		// To strictly test Punctuation, we should simulate the pipeline state.
		// However, to keep it simple and consistent with other tests, we will run the full pipeline
		// (Markers -> Punctuation) in the test loop to match the Golden Test expectations.

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

			// 4. Compare
			// We join the tokens to compare with the expected string,
			// because Punctuation logic affects spacing which is best seen in the joined string?
			// No, the Golden Tests define the output string.
			// But our functions return []string.
			// We should tokenize the expected string to compare slices,
			// OR we need a Joiner to compare final strings.
			// Since we don't have a Joiner yet (TASK-08), we will compare tokens.
			expectedTokens := functions.Tokenize(tt.expected)

			// Note: Tokenizing the expected string "word, (up)" gives ["word", ",", "(up)"]
			// because Tokenizer pads punctuation.
			// But ProcessPunctuation is supposed to MERGE punctuation to the word?
			// If ProcessPunctuation returns ["word,", "(up)"],
			// and Tokenize("word, (up)") returns ["word", ",", "(up)"], they won't match.
			//
			// CRITICAL: The Tokenizer splits punctuation.
			// The Punctuation Step is supposed to "Fix spacing".
			// Does it merge tokens?
			// PRD 4.3: "Standard Punctuation... must adhere to the previous word".
			// If it adheres, it becomes part of the token? ["word,"]?
			// OR does it just ensure there is no space when printed?
			// Architecture says "Data Structure: []string".
			// If we keep them as separate tokens, the Joiner needs to know not to put a space.
			// But the Joiner is simple (usually `strings.Join(tokens, " ")`).
			// Therefore, ProcessPunctuation MUST merge tokens.
			//
			// So, if ProcessPunctuation returns ["word,"],
			// and we compare with Tokenize("word,") -> ["word", ","], they fail.
			//
			// Solution: We cannot use Tokenize(expected) for verification here because Tokenize splits what Punctuation merges.
			// We must manually construct the expected slice for the test, OR implement a simple Joiner in the test.
			// I will implement a simple space-join in the test to verify the visual output matches the Golden Test string.

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
