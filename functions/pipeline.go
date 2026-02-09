package function

import "strings"

// Process takes the raw input string and runs it through the full pipeline.
func Process(input string) string {
	tokens := Tokenize(input)
	tokens = ProcessMarkers(tokens)
	tokens = ProcessGrammar(tokens)
	tokens = ProcessCasing(tokens)

	// Filter empty strings (placeholders for removed markers)
	var cleanTokens []string
	for _, t := range tokens {
		if t != "" {
			cleanTokens = append(cleanTokens, t)
		}
	}
	tokens = cleanTokens

	tokens = ProcessPunctuation(tokens)
	return strings.Join(tokens, " ")
}
