package function

import "strings"

// Tokenize splits the input string into tokens based on whitespace and punctuation rules.
func Tokenize(s string) []string {
	// 1. Protect groups by replacing them with temporary placeholders
	// Using unlikely strings as placeholders
	s = strings.ReplaceAll(s, "...", " __ELLIPSIS__ ")
	s = strings.ReplaceAll(s, "?!", " __QMARK_BANG__ ")
	s = strings.ReplaceAll(s, "!?", " __BANG_QMARK__ ")

	// 2. Pad standard punctuation
	punctuation := []string{".", ",", "!", "?", ":", ";"}
	for _, p := range punctuation {
		s = strings.ReplaceAll(s, p, " "+p+" ")
	}

	// 3. Restore protected groups
	// Note: We padded them with spaces in step 1 to ensure separation
	s = strings.ReplaceAll(s, "__ELLIPSIS__", "...")
	s = strings.ReplaceAll(s, "__QMARK_BANG__", "?!")
	s = strings.ReplaceAll(s, "__BANG_QMARK__", "!?")

	// 4. Split by whitespace
	return strings.Fields(s)
}
